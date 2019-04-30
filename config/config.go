package config

import (
	"apiproject/log"
	"github.com/Unknwon/goconfig"
	"go.uber.org/zap"
)

var Config *goconfig.ConfigFile

func init() {
	//加载配置文件
	var err error
	Config, err = goconfig.LoadConfigFile("config/config.ini")
	if err != nil {
		log.Logger.Error("加载配置文件异常", zap.String("error", err.Error()))
	}
}
