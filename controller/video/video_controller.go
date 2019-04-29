package video

import (
	"apiproject/service/video"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
查找视频列表接口
*/
func FindVideoList(ctx *gin.Context) {
	videoList := video.FindVideoList()
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": videoList,
	})
}
