package config

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

var Config *goconfig.ConfigFile

func init() {
	//加载配置文件
	var err error
	Config, err = goconfig.LoadConfigFile("config/config.ini")
	if err != nil {
		fmt.Println("加载配置文件异常[%v]", err)
	}
}
