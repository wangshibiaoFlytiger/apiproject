package video

import (
	"apiproject/service/video"
	"github.com/gin-gonic/gin"
	"net/http"
)

var videoService video.VideoService

func init() {
	//初始化依赖的service
	videoService = video.VideoService{}
}

/**
查找视频列表接口
*/
func FindVideoList(ctx *gin.Context) {
	videoList := videoService.FindVideoList()
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": videoList,
	})
}
