package main

import (
	"apiproject/bootstrap"
	"apiproject/config"
	"apiproject/cron"
	"apiproject/dao"
	"apiproject/router"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
)

func main() {
	//解析命令行参数
	parseCliParam()

	//系统初始化
	bootstrap.Init()

	defer dao.Db.Close()

	//初始化定时任务
	cron.Init()

	//初始化路由
	engine := router.Init()
	//如下代码放到最后, 否则其他代码没机会执行
	engine.Run(":8080")
}

/**
解析命令行参数
*/
func parseCliParam() {
	//实例化一个命令行程序
	app := cli.NewApp()
	//程序名称
	app.Name = "GoTool"
	//程序的用途描述
	app.Usage = "To save the world"
	//程序的版本号
	app.Version = "1.0.0"

	//设置启动参数
	app.Flags = []cli.Flag{
		//参数string, int, bool
		cli.StringFlag{
			Name:  "profile, p",        //参数名称
			Value: "dev",               //参数默认值
			Usage: "运行环境:dev,test,pro", //参数功能描述
		},
	}
	//该程序执行的代码
	app.Action = func(c *cli.Context) error {
		config.GlobalConfig.Profile = c.String("profile") //不使用变量接收，直接解析
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
