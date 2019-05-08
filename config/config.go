package config

import (
	"apiproject/env"
	"apiproject/log"
	"github.com/Unknwon/goconfig"
	"go.uber.org/zap"
)

var Config *goconfig.ConfigFile

func Init() {
	//加载配置文件
	var err error
	configFile := "config/config_" + env.SysEnv.Profile + ".ini"
	Config, err = goconfig.LoadConfigFile(configFile)
	if err != nil {
		log.Logger.Error("加载配置文件异常", zap.String("error", err.Error()))
	}
}
