package router

import (
	"apiproject/controller"
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

	//配置静态文件目录
	engine.Static("/static", "./public/static")
	//配置单个静态文件
	engine.StaticFile("/test.html", "./public/static/test.html")

	//配置模板路径
	engine.LoadHTMLGlob("./public/template/*")
	//配置首页入口
	engine.GET("/index", controller.Index)

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
