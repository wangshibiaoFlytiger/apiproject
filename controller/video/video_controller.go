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
	//绑定query参数到对象, 注意:默认情况下, 查询参数和对象成员的大小写需保持一致. 此外参数名可以通过如下tag自定义: form:"siteId" binding:"required"
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
	updateParamMap := make(map[string]interface{})
	updateParamMap["title"] = ctx.PostForm("title")
	updateParamMap["update_time"] = time.Now()

	//查询body数据
	bodyData, error := ctx.GetRawData()
	if error != nil {
		fmt.Printf("查询body数据异常[%v]", error)
	}
	fmt.Printf("body数据[%v]", string(bodyData))

	dao.Db.Model(&m_video.Video{}).Where("id = ?", "id2").Update(updateParamMap)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": nil,
	})
}
