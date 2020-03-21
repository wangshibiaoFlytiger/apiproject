package cron

import (
	"apiproject/config"
	"apiproject/log"
	s_task "apiproject/service/task"
	"github.com/robfig/cron/v3"
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

	cronSchduler := cron.New(cron.WithSeconds())

	err := registerTask(cronSchduler, "定时任务1", config.GlobalConfig.TaskTask1Cron, s_task.TaskService.Task1)
	if err != nil {
		panic(err)
	}
	cronSchduler.Start()

	err = registerTask(cronSchduler, "定时任务2", config.GlobalConfig.TaskTask1Cron, s_task.TaskService.Task2)
	if err != nil {
		panic(err)
	}
}

/**
注册定时任务
*/
func registerTask(cronSchduler *cron.Cron, jobName string, jobSpec string,
	callbackTask CallbackTask) error {
	log.Logger.Info("注册定时任务", zap.String("jobName", jobName), zap.String("jobSpec", jobSpec))

	//创建定时任务
	entryId, err := cronSchduler.AddFunc(
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

	log.Logger.Info("注册定时任务, 完成", zap.Any("entryId", entryId))

	return nil
}
