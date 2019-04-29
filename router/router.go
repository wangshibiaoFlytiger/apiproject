package router

import (
	"apiproject/controller/video"
	"github.com/gin-gonic/gin"
)

/**
初始化路由
*/
func Init() *gin.Engine {
	engine := gin.Default()

	engine.GET("/api/video/findList", video.FindVideoList)
	return engine
}
