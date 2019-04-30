package cron

import (
	"apiproject/log"
	"github.com/rfyiamcool/cronlib"
	"go.uber.org/zap"
)

func Init() {
	cronSchduler := cronlib.New()

	jobName := "我是任务1"
	jobSpec := "*/2 * * * * *"

	//创建定时任务
	job, err := cronlib.NewJobModel(
		jobSpec,
		func() {
			log.Logger.Info("创建定时任务", zap.String("jobSpec", jobSpec))
		},
	)
	if err != nil {
		panic(err.Error())
	}

	//注册定时任务
	err = cronSchduler.Register(jobName, job)
	if err != nil {
		panic(err.Error())
	}

	cronSchduler.Start()
}
