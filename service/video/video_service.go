package video

import (
	dao_video "apiproject/dao/video"
	"apiproject/model/video"
)

type VideoService struct {
}

var videoDao dao_video.VideoDao

func init() {
	//初始化依赖的dao
	videoDao = dao_video.VideoDao{}
}

//查询视频列表
func (this VideoService) FindVideoList() []video.VideoModel {
	videoList := videoDao.FindVideoList()
	return videoList
}
