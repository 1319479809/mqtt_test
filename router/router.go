package router

import (
	"fmt"
	"net/http"

	"github.com/1319479809/mqtt_test/device"
	"github.com/gin-gonic/gin"
)

// 门禁设备
func initDevice(r *gin.Engine) {
	v1 := r.Group("/device")
	{
		v1.POST("/control", device.DeviceControl) //远程控制设备
	}
	v2 := r.Group("/person")
	{
		v2.POST("/create", device.PersonCreate)               //人员注册.
		v2.POST("/delete", device.PersonDelete)               //人员删除
		v2.POST("/find", device.PersonFind)                   //人员查询
		v2.POST("/whiteListSync", device.PersonWhiteListSync) //同步白名单
		v2.POST("/whiteListFind", device.PersonWhiteListFind) //查询白名单
		v2.POST("/registerFeats", device.PersonRegisterFeats) //人员注册（feature）

	}
	v3 := r.Group("/v1")
	{
		v3.POST("/post", device.SendPost)
	}
	//定义默认路由
	r.NoRoute(func(c *gin.Context) {
		fmt.Println("test4")
		c.JSON(http.StatusNotFound, gin.H{
			"status": 0,
			"error":  "success",
		})
	})
	r.Run(":8086")
}
