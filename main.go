package main

import (
	"apiproject/router"
)

func main() {
	//初始化路由
	engine := router.Init()
	engine.Run(":8080")
}
