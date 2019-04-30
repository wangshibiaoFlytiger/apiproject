package router

import (
	c_video "apiproject/controller/video"
	"apiproject/middleware"
	"github.com/gin-gonic/gin"
)

/**
初始化路由
*/
func Init() *gin.Engine {
	engine := gin.Default()

	//视频相关接口
	videoGroup := engine.Group("/api/video", middleware.ApiLogMiddleware)
	videoGroup.GET("/findList", c_video.FindVideoList)
	videoGroup.GET("/findVideoByWhere", c_video.FindVideoByWhere)
	videoGroup.GET("/addVideo", c_video.AddVideo)
	videoGroup.GET("/updateVideo", c_video.UpdateVideo)
	return engine
}
