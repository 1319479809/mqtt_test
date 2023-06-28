package main

import (
	"fmt"
	"log"
	"math/rand"

	"time"

	"github.com/1319479809/mqtt_test/router"
	"github.com/1319479809/mqtt_test/test_test"
	"github.com/1319479809/mqtt_test/utils"
	"github.com/1319479809/mqtt_test/utils/slog"

	"github.com/gin-gonic/gin"
)

// http://192.168.2.19:80/sendTest

func main() {

	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)

	if utils.Cfg.Section("").Key("runmode").String() == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	slog.CpInfo("test", "test")
	test_test.GetRateUserValue(0, 1, 1)
	//getOutBoundTransferAdvisor()
	r := gin.Default()
	router.InitDeviceHttp(r)
	test_test.XmlTest()
	log.Println("end=====================")
}

// [{user_id:xx,Weights：xxx}]
func getOutBoundTransferAdvisor() {

	members := []string{"A", "B", "", "D", "E"}
	weights := []int{10, 8, 2, 5, 3}
	result := []string{}
	rand.Seed(time.Now().UnixNano())
	totalweight := 0

	for _, w := range weights {
		totalweight += w
	}

	for i := 0; i < 20; i++ {
		r := rand.Intn(totalweight)
		fmt.Println("r=", r, "totalweight=", totalweight)
		for j, w := range weights {
			if r < w {
				result = append(result, members[j])
				break
			}
			r -= w
		}

	}

	fmt.Println("result=========", result)
}
