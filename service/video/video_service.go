package s_video

import (
	"apiproject/dao"
	d_video "apiproject/dao/video"
	"apiproject/entity"
	"apiproject/log"
	"apiproject/model"
	m_video "apiproject/model/video"
	"apiproject/util"
	"go.uber.org/zap"
	"time"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

var VideoService = &videoService{}

type videoService struct {
}

//查询视频列表
func (this *videoService) FindVideoList() (videoList []*m_video.Video, err error) {
	if err := d_video.VideoDao.FindList(dao.Db, &videoList); err != nil {
		return nil, err
	}

	return videoList, nil
}

//分页查询视频列表
func (this *videoService) FindVideoListPage(pageNo int, pageSize int) *model.Page {
	page := &model.Page{
		PageNo:   pageNo,
		PageSize: pageSize,
		ItemList: &[]*m_video.Video{},
	}
	if err := d_video.VideoDao.FindPageData(dao.Db.Model(&m_video.Video{}), page); err != nil {
		return nil
	}

	return page
}

/**
批量添加视频
*/
func (this *videoService) BulkAddVideo() {
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

	error := d_video.VideoDao.BulkInsert(dao.Db, videoList, []string{"id", "site_id", "title", "created_at", "updated_at"})
	if error != nil {
		log.Logger.Error("批量添加视频失败", zap.Error(error))
	}
}
