package video

import (
	"apiproject/dao"
	"apiproject/model/video"
)

type VideoDao struct {
}

/**
查找视频列表
*/
func (this VideoDao) FindVideoList() []video.VideoModel {
	var videoList []video.VideoModel
	dao.Db.Find(&videoList)

	return videoList
}
