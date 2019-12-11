package cron

import (
	"apiproject/config"
	"apiproject/log"
	"github.com/rfyiamcool/cronlib"
	"go.uber.org/zap"
)

type CallbackTask func()

/**
启动定时任务列表
*/
func Init() {
	//任务开关是否开启
	if !config.GlobalConfig.TaskSwitch {
		log.Logger.Info("启动定时任务列表, 任务开关没有开启")
		return
	}

	cronSchduler := cronlib.New()

	err := registerTask(cronSchduler, "定时任务1", config.GlobalConfig.TaskTask1Cron, taskService.Task1)
	if err != nil {
		panic(err)
	}

	cronSchduler.Start()
}

/**
注册新闻爬取任务
*/
func registerTask(cronSchduler *cronlib.CronSchduler, jobName string, jobSpec string,
	callbackTask CallbackTask) error {
	log.Logger.Info("注册定时任务", zap.String("jobName", jobName), zap.String("jobSpec", jobSpec))

	//创建定时任务
	job, err := cronlib.NewJobModel(
		jobSpec,
		func() {
			log.Logger.Info("定时任务开始执行", zap.Any("jobName", jobName), zap.Any("jobSpec", jobSpec))
			callbackTask()
			log.Logger.Info("定时任务执行完成", zap.Any("jobName", jobName), zap.Any("jobSpec", jobSpec))
		},
	)
	if err != nil {
		return err
	}

	//注册定时任务
	err = cronSchduler.Register(jobName, job)
	if err != nil {
		return err
	}

	return nil
}
