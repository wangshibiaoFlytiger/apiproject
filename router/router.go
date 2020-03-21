package router

import (
	apiproject_config "apiproject/config"
	"apiproject/controller"
	c_kafka "apiproject/controller/kafka"
	c_video "apiproject/controller/video"
	c_wxpay "apiproject/controller/wxpay"
	"apiproject/log"
	"apiproject/middleware"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

/**
初始化路由
*/
func Init() *gin.Engine {
	engine := gin.Default()

	/**************************start 配置中间件 *****************/
	//支持跨域
	config := cors.DefaultConfig()
	config.AddAllowHeaders("X-Requested-With")
	config.AllowAllOrigins = true
	engine.Use(cors.New(config))
	//全局配置api日志访问中间件
	engine.Use(middleware.ApiLogMiddleware)
	/**************************end 配置中间件 *****************/

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
