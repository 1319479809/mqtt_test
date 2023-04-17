package main

// 与后端mqtt服务交互

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const EMQServerAddress = "你的IP"

// 创建全局mqtt publish消息处理 handler
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("Push Message:")
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

// 创建全局mqtt sub消息处理 handler
var messageSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("收到订阅消息:")
	fmt.Printf("Sub Client Topic : %s \n", msg.Topic())
	fmt.Printf("Sub Client msg : %s \n", msg.Payload())
}

// 连接的回掉函数
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("新的连接!" + " Connected")
}

// 丢失连接的回掉函数
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect loss: %v\n", err)
}

func init() {
	// 配置错误提示
	mqtt.DEBUG = log.New(os.Stdout, "		[mqttDEBUG]", 0)
	mqtt.ERROR = log.New(os.Stdout, " 	[mqttERROR]", 0)
}

/**
 * @Description: 发布订阅
 * @param clientID
 * @param addr
 * @param topic
 * @param payload
 */
func Push(topic string, qos byte, retain bool, payload string) {
	// opts ClientOptions 用于设置 broker，端口，客户端 id ，用户名密码等选项
	opts := mqtt.NewClientOptions().AddBroker("tcp://" + EMQServerAddress + ":1883").SetClientID("test_push")
	opts.SetKeepAlive(60 * time.Second)
	// Message callback handler，在没有任何订阅时，发布端调用此函数
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.SetPingTimeout(1 * time.Second)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	//发布消息
	// qos是服务质量: ==1: 一次, >=1: 至少一次, <=1:最多一次
	// retained: 表示mqtt服务器要保留这次推送的信息，如果有新的订阅者出现，就会把这消息推送给它（持久化推送）
	token := client.Publish(topic, qos, retain, payload)
	token.Wait()
	fmt.Println("Push Data : "+topic, "Data Size is "+strconv.Itoa(len(payload)))
	fmt.Println("Disconnect with broker")
	client.Disconnect(250)
}

/**
 * @Description: 订阅与取消订阅
 * @param clientID
 * @param addr
 * @param topic
 * @param isSub
 */
func Subscription(topic string, qos byte, isSub bool, handleFun func([]byte)) {
	opts := mqtt.NewClientOptions().AddBroker("tcp://" + EMQServerAddress + ":1883").SetClientID("sub_test")
	opts.SetKeepAlive(60 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	opts.OnConnect = func(client mqtt.Client) {
		fmt.Println("New Subscription! Connected" + " => " + topic)
	}
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if isSub {
		// 订阅消息
		if token := client.Subscribe(topic, qos, func(client mqtt.Client, msg mqtt.Message) {
			fmt.Printf("Receive Subscribe Message :")
			fmt.Printf("Sub Client Topic : %s, Data size is  %d \n", msg.Topic(), len(msg.Payload()))
			if len(msg.Payload()) > 0 {
				handleFun(msg.Payload())
			}
		}); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
	} else {
		// 取消订阅
		if token := client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
	}
}

func main() {
	fmt.Println("123")
}
