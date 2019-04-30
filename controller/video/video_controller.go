package c_video

import (
	"apiproject/dao"
	m_video "apiproject/model/video"
	s_video "apiproject/service/video"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
	siteId := ctx.Query("siteId")

	videoList := []m_video.Video{}
	dao.Db.Where("site_id = ? and title like ?", siteId, "%7%").Find(&videoList)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": videoList,
	})
}

/**
添加视频
*/
func AddVideo(ctx *gin.Context) {
	video := m_video.Video{}
	video.ID = "id2"
	video.Title = "title2"
	now := time.Now()
	video.CreateTime = now
	video.UpdateTime = now
	dao.Db.Create(video)
}

/**
更新视频
*/
func UpdateVideo(ctx *gin.Context) {
	updateParamMap := make(map[string]interface{})
	updateParamMap["title"] = "正在播放:韩国美女激情VIP秀1071_update2019-04-30"
	updateParamMap["update_time"] = time.Now()
	dao.Db.Model(&m_video.Video{}).Where("id = ?", "id2").Update(updateParamMap)
}
