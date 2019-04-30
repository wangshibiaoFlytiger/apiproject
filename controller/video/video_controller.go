package c_video

import (
	"apiproject/dao"
	m_video "apiproject/model/video"
	s_video "apiproject/service/video"
	"github.com/gin-gonic/gin"
	"net/http"
)

var videoService s_video.VideoService

func init() {
	//初始化依赖的service
	videoService = s_video.VideoService{}
}

/**
查询视频列表接口
*/
func FindVideoList(ctx *gin.Context) {
	videoList := videoService.FindVideoList()
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": videoList,
	})
}

/**
按条件查询视频列表接口
*/
func FindVideoByWhere(ctx *gin.Context) {
	videoList := []m_video.Video{}
	dao.Db.Where("site_id = ? and title like ?", "nvrenwu", "%7%").Find(&videoList)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": videoList,
	})
}

/**
添加视频
*/
func AddVideo(ctx *gin.Context) {
	video := m_video.Video{"id1", "title1"}
	dao.Db.Create(video)
}

/**
更新视频
*/
func UpdateVideo(ctx *gin.Context) {
	updateParamMap := make(map[string]string)
	updateParamMap["title"] = "正在播放:韩国美女激情VIP秀1071_update2019-04-30"
	dao.Db.Model(&m_video.Video{}).Where("id = ?", "1c29614d6274bf4566c7816581eef9e8").Update(updateParamMap)
}
