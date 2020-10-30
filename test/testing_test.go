package test

import (
	s_video "apiproject/service/video"
	"testing"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

/**
测试service层方法
*/
func TestVideoService(t *testing.T) {
	videoList, _ := s_video.VideoService.FindVideoList()
	t.Log(videoList)
}
