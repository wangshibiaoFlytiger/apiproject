package main

import (
	"apiproject/bootstrap"
	"apiproject/config"
	"apiproject/dao"
	"apiproject/log"
	"apiproject/router"
	"apiproject/util"
	"go.uber.org/zap"
	"gopkg.in/urfave/cli.v1"
	"os"
)

//当有新的.go源文件引用静态文件时, 注意同步修改如下命令的--import-path参数
//go:generate rice -v --import-path "./router" --import-path "./config"  embed-go

//生成swagger api接口文档, 注意指定-g参数
//go:generate swag init -g router/router.go
func main() {
	//解析命令行参数
	parseCliParam()

	//系统初始化
	bootstrap.Init()

	defer dao.Db.Close()

	//初始化路由
	engine := router.Init()
	//如下代码放到最后, 否则其他代码没机会执行
	if err := engine.Run(":" + util.IntToStr(config.GlobalConfig.ServicePort)); err != nil {
		log.Logger.Error("main函数异常", zap.Error(err))
	}
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
			Name:  "profile, p",                                   //参数名称
			Value: "dev",                                          //参数默认值
			Usage: "运行环境:back-dev,back-test,front-dev,front-test", //参数功能描述
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
