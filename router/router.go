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
	engine.GET("/api/video/findVideoByWhere", video.FindVideoByWhere)
	engine.GET("/api/video/addVideo", video.AddVideo)
	engine.GET("/api/video/updateVideo", video.UpdateVideo)
	return engine
}
