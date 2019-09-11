package bootstrap

import (
	"apiproject/cache"
	"apiproject/config"
	"apiproject/cron"
	"apiproject/dao"
	"apiproject/ip_location"
	"apiproject/kafka"
	"apiproject/log"
	"apiproject/mongo"
)

/**
系统初始化
*/
func Init() {
	config.Init()
	log.Init()

	ip_location.Init()

	cache.Init()
	kafka.Init()
	mongo.Init()

	//dao层初始化
	dao.Init()

	//初始化定时任务
	cron.Init()
}
