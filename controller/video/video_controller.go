package c_video

import (
	"apiproject/dao"
	m_video "apiproject/model/video"
	s_video "apiproject/service/video"
	"fmt"
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
	/**
	绑定query参数到对象.
	注意:默认情况下, "query参数或form参数"和对象成员的大小写需保持一致. 除非为对象补充tag定义, 定义规则如下:
	1. 对于query参数, 需要如下定义
	form:"siteId" binding:"required"
	2. 对于form参数, 需要如下定义
	form:"siteId"
	3. 对于body的json参数, 无需定义tag, 且不区分大小写
	*/
	videoQuery := m_video.Video{}
	if ctx.Bind(&videoQuery) == nil {
		fmt.Printf("绑定后的videoQuery[%v]", videoQuery)
	}

	videoList := []m_video.Video{}
	dao.Db.Where("site_id = ? and title like ?", videoQuery.SiteId, "%7%").Find(&videoList)
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

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": nil,
	})
}

/**
更新视频
*/
func UpdateVideo(ctx *gin.Context) {
	//绑定参数到对象
	videoBind := m_video.Video{}
	if ctx.Bind(&videoBind) == nil {
		fmt.Printf("绑定后的videoBind[%v]", videoBind)
	}
	videoBind.UpdateTime = time.Now()

	dao.Db.Model(&m_video.Video{}).Where("id = ?", "id2").Update(videoBind)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": nil,
	})
}
