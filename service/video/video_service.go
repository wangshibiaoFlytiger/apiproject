package s_video

import (
	m_video "apiproject/model/video"
)

type VideoService struct {
}

//查询视频列表
func (this VideoService) FindVideoList() []m_video.Video {
	videoList := videoDao.FindVideoList()
	return videoList
}
