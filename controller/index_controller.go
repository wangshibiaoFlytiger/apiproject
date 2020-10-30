package controller

import (
	"apiproject/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

//渲染首页入口页面
func Index(ctx *gin.Context) {
	serviceName := config.GlobalConfig.ServiceName
	//配置前端服务页面
	if serviceName == "apiprojectBack" {
		ctx.HTML(http.StatusOK, "back.html", gin.H{
			"title":   "我是后台管理",
			"encrypt": config.GlobalConfig.ServiceApiResponseEncrypt,
		})
	} else if serviceName == "apiprojectFront" { //配置后台管理页面
		ctx.HTML(http.StatusOK, "front.html", gin.H{
			"title":   "我是前端页面",
			"encrypt": config.GlobalConfig.ServiceApiResponseEncrypt,
		})
	}
}
