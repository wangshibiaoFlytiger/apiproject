package test

import (
	s_video "apiproject/service/video"
	"testing"
)

/**
测试service层方法
*/
func TestVideoService(t *testing.T) {
	videoList := s_video.VideoService.FindVideoList()
	t.Log(videoList)
}
