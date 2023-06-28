package utils

import (
	"encoding/xml"
	"errors"
	"log"
	"strings"
)

type XRegisterBetaMiniprogram struct {
	CommentAppID string `xml:"AppId"` // 第三方平台appid
	CreateTime   string `xml:"CreateTime"`
	InfoType     string `xml:"InfoType"`
	AppID        string `xml:"appid"`
	Status       int    `xml:"status"`
	Msg          string `xml:"msg"`
	Info         struct {
		Name     string `xml:"name"`
		UniqueId string `xml:"unique_id"`
	} `xml:"info"`
}

type XRegisterBetaMiniprogram2 struct {
	CommentAppID string `xml:"AppId"` // 第三方平台appid
	CreateTime   string `xml:"CreateTime"`
	InfoType     string `xml:"InfoType"`
	AppID        string `xml:"appid"`
	Status       int    `xml:"status"`
	Msg          string `xml:"msg"`
	Info         struct {
		Name               string `xml:"name"`
		Code               string `xml:"code"`
		CodeType           int    `xml:"code_type"`
		LegalPersonaWechat string `xml:"legal_persona_wechat"`
		LegalPersonaName   string `xml:"legal_persona_name"`
		ComponentPhone     string `xml:"component_phone"`
	} `xml:"info"`
}

type XRegister struct {
	CommentAppID string      `xml:"AppId"` // 第三方平台appid
	CreateTime   string      `xml:"CreateTime"`
	InfoType     string      `xml:"InfoType"`
	AppID        string      `xml:"appid"`
	Status       int         `xml:"status"`
	Msg          string      `xml:"msg"`
	Info         interface{} `xml:"info"`
}

type Info1 struct {
	Name     string `xml:"name"`
	UniqueId string `xml:"unique_id"`
}

type Info2 struct {
	Name               string `xml:"name"`
	Code               string `xml:"code"`
	CodeType           int    `xml:"code_type"`
	LegalPersonaWechat string `xml:"legal_persona_wechat"`
	LegalPersonaName   string `xml:"legal_persona_name"`
	ComponentPhone     string `xml:"component_phone"`
}

func TestEvent(result string) (event *XRegisterBetaMiniprogram, err error) {
	s := strings.Index(result, "<")
	e := strings.LastIndex(result, ">")
	if s == -1 || e == -1 {
		return nil, errors.New("Msg Font Error")
	}

	log.Println(result[s : e+1])
	event = new(XRegisterBetaMiniprogram)
	err = xml.Unmarshal([]byte(result[s:e+1]), event)
	if err != nil {
		return nil, errors.New("xml Unmarshal Error")
	}
	log.Println(event)
	return event, nil
}
