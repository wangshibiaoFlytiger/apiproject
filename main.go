package main

import (
	"apiproject/cron"
	"apiproject/dao"
	"apiproject/router"
)

func main() {
	defer dao.Db.Close()

	//初始化定时任务
	cron.Init()

	//初始化路由
	engine := router.Init()
	//如下代码放到最后, 否则其他代码没机会执行
	engine.Run(":8080")
}
