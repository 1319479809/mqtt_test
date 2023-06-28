package utils

import (
	"github.com/1319479809/mqtt_test/utils/slog"
	"gopkg.in/ini.v1"
)

// 系统配置参数
var (
	Cfg *ini.File
)

func init() {
	var err error
	Cfg, err = LoadConf("conf/app.conf")
	if err != nil {
		slog.Error("init", err)
	}

	go initIdGenerator()
}

// LoadConf 加载配置文件
func LoadConf(file string) (cfg *ini.File, err error) {
	cfg, err = ini.Load(file)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// GetRunHTTPPort 获取Http端口
func GetRunHTTPPort() (port string) {
	port = Cfg.Section("").Key("httpport").String()
	if port == "" {
		return "8080"
	}
	return port
}
