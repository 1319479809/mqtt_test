package test_test

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/1319479809/mqtt_test/utils"
)

// 解析xml
func XmlTest() {
	file, err := os.Open("test.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	var adata0 utils.XRegister
	var info1 utils.Info1
	adata0.Info = info1
	err = xml.Unmarshal(data, &adata0)
	if err != nil {
		log.Println("err=", err)
	}
	log.Println("adata01=", adata0)

	var adata02 utils.XRegister
	var info2 utils.Info2
	adata02.Info = info2
	err = xml.Unmarshal(data, &adata02)
	if err != nil {
		log.Println("err=", err)
	}
	log.Println("adata02=", adata02)
	//info := utils.XRegister.Info

	var adata utils.XRegisterBetaMiniprogram
	err = xml.Unmarshal(data, &adata)
	if err != nil {
		log.Println("err=", err)
	}
	info := adata.Info
	log.Println("info=", info)
	log.Println("info=", info.UniqueId)
	log.Println("adata=", adata)

	var adata2 utils.XRegisterBetaMiniprogram2
	err = xml.Unmarshal(data, &adata2)
	if err != nil {
		log.Println("err=", err)
	}
	log.Println("adata2=", adata2)

	// res, err := utils.TestEvent(string(data))
	// if err != nil {
	// 	log.Println("err=", err)
	// }
	// log.Println("res=", res)
}
