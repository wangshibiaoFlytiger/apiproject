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
	s_cron "apiproject/service/cron"
	"go.uber.org/zap"
)

/**
系统初始化
*/
func Init() {
	config.Init()

	cache.Init()
	kafka.Init()
	mongo.Init()
	//dao层初始化
	dao.Init()

	log.Init()

	ip_location.Init()

	//初始化定时任务
	cron.Init()
	if err := s_cron.CronTaskService.LoadCronTaskList(); err != nil {
		log.Logger.Error("系统初始化, 加载定时任务列表, 异常", zap.Error(err))
	}
}
