package main

import (
	"apiproject/dao"
	"apiproject/router"
)

func main() {
	//初始化路由
	engine := router.Init()
	engine.Run(":8080")

	defer dao.Db.Close()
}
