package test

import (
	s_video "apiproject/service/video"
	"testing"
)

/**
测试service层方法
*/
func TestVideoService(t *testing.T) {
	videoService := s_video.VideoService{}
	videoList := videoService.FindVideoList()
	t.Log(videoList)
}
