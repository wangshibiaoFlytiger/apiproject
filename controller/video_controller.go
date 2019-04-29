package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
查找视频列表接口
*/
func FindVideoList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": "这是业务数据",
	})
}
