package controller

import (
	"apiproject/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

//渲染首页入口页面
func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "我是模板变量的内容",
		"encrypt": config.GlobalConfig.ServiceApiResponseEncrypt,
	})
}
