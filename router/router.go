package router

import (
	"apiproject/controller"
	c_video "apiproject/controller/video"
	"apiproject/middleware"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"html/template"
)

/**
初始化路由
*/
func Init() *gin.Engine {
	engine := gin.Default()

	//通过go.rice配置静态文件目录
	engine.StaticFS("/static", rice.MustFindBox("../public/static").HTTPBox())

	/***********************start 通过go.rice配置页面模板 **********************/
	//配置模板文件的根目录
	templateBox := rice.MustFindBox("../public/template")

	//配置模板文件路径列表, 需填写相对于模板相对路径
	list := [...]string{"index.html"}
	for _, x := range list {
		templateString, err := templateBox.String(x)
		if err != nil {
			panic(err)
		}

		tmplMessage, err := template.New(x).Parse(templateString)
		if err != nil {
			panic(err)
		}

		engine.SetHTMLTemplate(tmplMessage)
	}
	/***********************end 配置页面模板 **********************/

	//配置首页入口
	engine.GET("/", controller.Index)
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
