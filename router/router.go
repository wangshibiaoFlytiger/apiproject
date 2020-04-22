package router

import (
	apiproject_config "apiproject/config"
	"apiproject/controller"
	c_cron "apiproject/controller/cron"
	c_kafka "apiproject/controller/kafka"
	c_video "apiproject/controller/video"
	c_wxpay "apiproject/controller/wxpay"
	"apiproject/docs"
	"apiproject/log"
	"apiproject/middleware"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

/**
初始化路由
*/

/************************start swagger api定义注解 **************/
// @contact.name Wang Shibiao
// @contact.email 645102170@qq.com
/************************end swagger api定义注解 **************/
func Init() *gin.Engine {
	engine := gin.Default()

	/**************************start cors跨域中间件 *****************/
	//支持跨域
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowHeaders("X-Requested-With")
	corsConfig.AllowAllOrigins = true
	engine.Use(cors.New(corsConfig))
	/**************************end cors跨域中间件 *****************/

	/**************************start 访问日志中间件 *****************/
	//全局配置api日志访问中间件
	engine.Use(middleware.ApiLogMiddleware)
	/**************************end 访问日志中间件 *****************/

	/**************************start gzip数据压缩中间件 *****************/
	/**
	客户端若想获取压缩的数据, 则需要指定http头, 测试过程如下
	1. 不指定header头, 请求返回的结果没有经过压缩
	curl http://localhost:8080/api/cronTask/findCronTaskList -o ./api_resp_src
	2. 接收压缩的数据, 同时保存压缩数据到本地文件
	curl http://localhost:8080/api/cronTask/findCronTaskList -H 'Accept-Encoding: gzip, deflate, sdch' -o ./api_resp_compress.gz
	保存的本地gip格式的压缩文件, 可以使用如下命令解压: gzip -d ./api_resp_compress.gz
	3. 使用--compressed将会自动解压收到的压缩数据, 保存解压后的数据到本地文件
	curl http://localhost:8080/api/cronTask/findCronTaskList -H 'Accept-Encoding: gzip, deflate, sdch' --compressed -o ./api_resp_auto_uncompress
	*/
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	/**************************end gzip数据压缩中间件 *****************/

	/***********************start 通过go.rice配置页面模板 **********************/
	//配置模板文件的根目录
	templateBox := rice.MustFindBox("../public/template")
	renderer := multitemplate.NewRenderer()
	//配置模板文件路径列表, 需填写相对于模板相对路径
	fileNameList := [...]string{"front.html", "back.html"}
	for _, fileName := range fileNameList {
		templateString, err := templateBox.String(fileName)
		if err != nil {
			panic(err)
		}

		renderer.AddFromString(fileName, templateString)
	}
	engine.HTMLRender = renderer
	/***********************end 配置页面模板 **********************/

	//通过go.rice配置静态文件目录
	engine.StaticFS("/static", rice.MustFindBox("../public/static").HTTPBox())

	/***********************start swagger api接口文档相关 **********************/
	//自定义swagger相关变量
	docs.SwaggerInfo.Title = "API接口文档"
	docs.SwaggerInfo.Description = "这是golang开发的api服务"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	/***********************end swagger api接口文档相关 **********************/

	/***********************start 配置首页入口 **********************/
	engine.GET("/", controller.Index)
	engine.GET("/index", controller.Index)
	/***********************end 配置首页入口 **********************/

	/***********************start 视频相关接口 **********************/
	videoGroup := engine.Group("/api/video")
	//加密api响应数据中间件
	if apiproject_config.GlobalConfig.ServiceApiResponseEncrypt {
		videoGroup.Use(middleware.EncryptResponseMiddleware)
	}
	videoGroup.GET("/findList", c_video.FindVideoList)
	videoGroup.GET("/findVideoListPage", c_video.FindVideoListPage)
	videoGroup.GET("/findVideoByWhere", c_video.FindVideoByWhere)
	videoGroup.GET("/addVideo", c_video.AddVideo)
	videoGroup.GET("/bulkAddVideo", c_video.BulkAddVideo)
	videoGroup.POST("/updateVideo", c_video.UpdateVideo)
	videoGroup.POST("/deleteVideo", c_video.DeleteVideo)
	/***********************end 视频相关接口 **********************/

	/***********************start kafka相关接口 **********************/
	kafkaGroup := engine.Group("/api/kafka")
	//加密api响应数据中间件
	if apiproject_config.GlobalConfig.ServiceApiResponseEncrypt {
		kafkaGroup.Use(middleware.EncryptResponseMiddleware)
	}
	kafkaGroup.GET("/sendMessage", c_kafka.SendMessage)
	/***********************end kafka相关接口 **********************/

	/***********************start 微信支付相关接口 **********************/
	wxpayGroup := engine.Group("/api/wxpay")
	//加密api响应数据中间件
	if apiproject_config.GlobalConfig.ServiceApiResponseEncrypt {
		wxpayGroup.Use(middleware.EncryptResponseMiddleware)
	}
	wxpayGroup.POST("/wxH5Pay", c_wxpay.WxH5Pay)
	wxpayGroup.POST("/wxH5PayCallback", c_wxpay.WxH5PayCallback)
	/***********************end 微信支付相关接口 **********************/

	/***********************start 定时任务相关接口 **********************/
	cronTaskGroup := engine.Group("/api/cronTask")
	cronTaskGroup.POST("/addCronTask", c_cron.AddCronTask)
	cronTaskGroup.DELETE("/deleteCronTask", c_cron.DeleteCronTask)
	cronTaskGroup.POST("/enableCronTask", c_cron.EnableCronTask)
	cronTaskGroup.POST("/disableCronTask", c_cron.DisableCronTask)
	cronTaskGroup.GET("/findCronTaskList", c_cron.FindCronTaskList)
	/***********************end 定时任务相关接口 **********************/

	/***********************start 反向代理相关接口 **********************/
	reverseproxyList := []map[string]string{}
	jsoniter.UnmarshalFromString(apiproject_config.GlobalConfig.ReverseproxyList, &reverseproxyList)
	log.Logger.Info("初始化路由, 查询反向代理配置", zap.Any("reverseproxyList", reverseproxyList))
	for _, reverseProxy := range reverseproxyList {
		//代理匹配urlPrefix的api到target服务
		engine.Use(middleware.ReverseProxyMiddleware(reverseProxy["urlPrefix"], middleware.ProxyOption{
			Target:      reverseProxy["target"],
			PathRewrite: "",
		}))
	}
	/***********************end 反向代理相关接口 **********************/

	return engine
}
