package test

import (
	s_video "apiproject/service/video"
	"testing"
)

/**
测试service层方法
*/
func TestVideoService(t *testing.T) {
	videoList, _ := s_video.VideoService.FindVideoList()
	t.Log(videoList)
}
