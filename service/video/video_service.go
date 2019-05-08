package s_video

import (
	d_video "apiproject/dao/video"
	m_video "apiproject/model/video"
)

type VideoService struct {
}

var videoDao d_video.VideoDao

func Init() {
	//初始化依赖的dao
	videoDao = d_video.VideoDao{}
}

//查询视频列表
func (this VideoService) FindVideoList() []m_video.Video {
	videoList := videoDao.FindVideoList()
	return videoList
}
