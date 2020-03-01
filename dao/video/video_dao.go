package d_video

import (
	"apiproject/dao"
	"apiproject/model"
	m_video "apiproject/model/video"
)

var VideoDao = &videoDao{}

type videoDao struct {
	dao.BaseDao
}

/**
查找视频列表
*/
func (this *videoDao) FindVideoList() []m_video.Video {
	var videoList []m_video.Video
	dao.Db.Find(&videoList)

	return videoList
}

/**
分页查找视频列表
*/
func (this *videoDao) FindVideoListPage(pageNo int, pageSize int) *model.Page {
	db := dao.Db.Model(&m_video.Video{})

	page := &model.Page{
		PageNo:    pageNo,
		PageSize:  pageSize,
		PageCount: 0,
		ItemCount: 0,
		ItemList:  &[]m_video.Video{},
	}
	if err := this.FindPageData(db, page); err != nil {
		return nil
	}

	return page
}
