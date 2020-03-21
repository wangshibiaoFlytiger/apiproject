package cron

import (
	"apiproject/log"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

//全局定时任务调度器
var CronSchduler *cron.Cron

type CallbackTask func()

/**
启动定时任务列表
*/
func Init() {
	//启动全局定时任务调度器
	CronSchduler = cron.New(cron.WithSeconds())
	CronSchduler.Start()
}

/**
注册定时任务
*/
func RegisterTask(jobName string, jobSpec string, callbackTask CallbackTask) (entryId cron.EntryID, err error) {
	log.Logger.Info("注册定时任务", zap.String("jobName", jobName), zap.String("jobSpec", jobSpec))

	//创建定时任务
	entryId, err = CronSchduler.AddFunc(
		jobSpec,
		func() {
			log.Logger.Info("定时任务开始执行", zap.Any("jobName", jobName), zap.Any("jobSpec", jobSpec))
			callbackTask()
			log.Logger.Info("定时任务执行完成", zap.Any("jobName", jobName), zap.Any("jobSpec", jobSpec))
		},
	)
	if err != nil {
		return 0, err
	}

	log.Logger.Info("注册定时任务, 完成", zap.Any("entryId", entryId))

	return entryId, nil
}

/**
反注册定时任务
*/
func UnRegisterTask(entryId cron.EntryID) {
	entry := CronSchduler.Entry(entryId)
	CronSchduler.Remove(entryId)
	log.Logger.Info("反注册定时任务, 完成", zap.Any("entry", entry))
}

/**
是否存在指定定时任务
*/
func ExistTask(entryId cron.EntryID) (exist bool) {
	entry := CronSchduler.Entry(entryId)
	if entry.Valid() {
		log.Logger.Info("是否存在指定定时任务, 是", zap.Any("entry", entry))
		return true
	}

	return false
}
