package router

import (
	c_video "apiproject/controller/video"
	"apiproject/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/**
初始化路由
*/
func Init() *gin.Engine {
	engine := gin.Default()

	//支持跨域
	engine.Use(cors.Default())

	//视频相关接口
	videoGroup := engine.Group("/api/video")
	//配置api日志访问中间件
	videoGroup.Use(middleware.ApiLogMiddleware)

	videoGroup.GET("/findList", c_video.FindVideoList)
	videoGroup.GET("/findVideoByWhere", c_video.FindVideoByWhere)
	videoGroup.GET("/addVideo", c_video.AddVideo)
	videoGroup.POST("/updateVideo", c_video.UpdateVideo)
	videoGroup.POST("/deleteVideo", c_video.DeleteVideo)
	return engine
}
