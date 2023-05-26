package main

import (
	"fmt"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// 12847(mqtt), 12173(mqtts), 8083(ws), 8084(wss)
const broker = "tcp://127.0.0.1:1883"
const username = ""
const password = ""
const ClientID = "go_mqtt1_client"

// message的回调
var onMessage mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("[%s] -> %s\n", msg.Topic(), msg.Payload())
}

var wg sync.WaitGroup
var client mqtt.Client

func main() {
	//连接MQTT服务器
	mqttConnect()
	defer client.Disconnect(250) //注册销毁
	wg.Add(1)
	go mqttSubScribe("golang-mqtt/test")
	// wg.Add(1)
	// go testPublish()
	wg.Wait()
}
func mqttConnect() {
	//配置
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(ClientID)
	// opts.SetUsername("")
	// opts.SetPassword("")
	opts.SetConnectTimeout(time.Duration(60) * time.Second)
	//连接
	client = mqtt.NewClient(opts)
	//客户端连接判断
	if token := client.Connect(); token.WaitTimeout(time.Duration(60)*time.Second) && token.Wait() && token.Error() != nil {
		fmt.Println("订阅 MQTT 失败")
		panic(token.Error())
	}
}

func mqttSubScribe(topic string) {
	defer wg.Done()
	for {
		// now := time.Now()
		// fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
		token := client.Subscribe(topic, 1, onMessage)
		token.Wait()
	}
}

// 测试 3秒发送一次，然后自己接收
func testPublish() {
	defer wg.Done()
	for {
		now := time.Now()
		//fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
		client.Publish("topic/test", 1, false, now.Format("2006-01-02 15:04:05.000 Mon Jan"))
		time.Sleep(time.Duration(3) * time.Second)
	}
}
