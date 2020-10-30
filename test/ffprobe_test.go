package test

import (
	"apiproject/util"
	"log"
	"testing"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

/**
测试ffprobe
*/
func TestFfprobe(t *testing.T) {
	videoPath := "/home/wangshibiao/test/test11.mp4"
	log.Println(util.GetVideoInfo(videoPath))
}
