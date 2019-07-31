package test

import (
	"apiproject/util"
	"log"
	"testing"
)

/**
测试ffprobe
*/
func TestFfprobe(t *testing.T) {
	videoPath := "/home/wangshibiao/test/test11.mp4"
	log.Println(util.GetVideoInfo(videoPath))
}
