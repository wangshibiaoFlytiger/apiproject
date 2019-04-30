package router

import (
	c_video "apiproject/controller/video"
	"github.com/gin-gonic/gin"
)

/**
初始化路由
*/
func Init() *gin.Engine {
	engine := gin.Default()

	engine.GET("/api/video/findList", c_video.FindVideoList)
	engine.GET("/api/video/findVideoByWhere", c_video.FindVideoByWhere)
	engine.GET("/api/video/addVideo", c_video.AddVideo)
	engine.GET("/api/video/updateVideo", c_video.UpdateVideo)
	return engine
}
