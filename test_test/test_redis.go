package test_test

import (
	"log"
	"time"

	"github.com/1319479809/mqtt_test/utils/auth"
)

// TestRedis 测试redis

func TestRedis() {
	str := "123456"
	log.Println("str=====", str)
	OneMinute := 1 * time.Minute
	err := auth.SetCache(str, "654321", OneMinute)
	if err != nil {
		log.Println("error: ", err)
	}
	time.Sleep(1 * time.Second)
	res, err := auth.GetCache(str)
	if err != nil {
		log.Println("error: ", err)
	}
	log.Println("res=", res)
}
