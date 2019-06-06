package bootstrap

import (
	"apiproject/cache"
	"apiproject/config"
	"apiproject/cron"
	"apiproject/dao"
	"apiproject/log"
)

/**
系统初始化
*/
func Init() {
	config.Init()
	log.Init()

	cache.Init()

	//dao层初始化
	dao.Init()

	//初始化定时任务
	cron.Init()
}
