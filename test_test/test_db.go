package test_test

import (
	"log"

	"github.com/1319479809/mqtt_test/utils/db"
)

// 测试db
func TestDb() {
	// TODO
	db := db.GetSystemOrmDb()
	BusinessId := 1001
	res, err := db.FindCreatMqttUserData(int64(BusinessId))
	if err != nil {
		log.Println("error: ", err)
	}
	log.Println("res=", res)
}
