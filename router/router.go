package router

import (
	"apiproject/controller"
	"github.com/gin-gonic/gin"
)

/**
初始化路由
*/
func Init() *gin.Engine {
	engine := gin.Default()

	engine.GET("/api/video/findList", controller.FindVideoList)
	return engine
}
