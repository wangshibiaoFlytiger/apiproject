package s_video

import (
	"apiproject/entity"
	"apiproject/log"
	m_video "apiproject/model/video"
	"apiproject/util"
	"go.uber.org/zap"
	"time"
)

type VideoService struct {
}

//查询视频列表
func (this *VideoService) FindVideoList() []m_video.Video {
	videoList := videoDao.FindVideoList()
	return videoList
}

/**
批量添加视频
*/
func (this *VideoService) BulkAddVideo() {
	jsonTime := &entity.JsonTime{time.Now()}

	video1 := m_video.Video{}
	video1.ID = util.GenUniqueId()
	video1.Title = "video1"
	video1.CreatedAt = jsonTime
	video1.UpdatedAt = jsonTime

	video2 := m_video.Video{}
	video2.ID = util.GenUniqueId()
	video2.Title = "video2"
	video2.CreatedAt = jsonTime
	video2.UpdatedAt = jsonTime

	videoList := []interface{}{}
	videoList = append(videoList, video1)
	videoList = append(videoList, video2)

	error := videoDao.BulkInsert(videoList, []string{"id", "site_id", "title", "created_at", "updated_at"})
	if error != nil {
		log.Logger.Error("批量添加视频失败", zap.Error(error))
	}
}
