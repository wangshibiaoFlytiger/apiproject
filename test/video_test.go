package test

import (
	"apiproject/util"
	"testing"
)

/**
测试视频转码
*/
func TestTranscodeVideo(t *testing.T) {
	err := util.TranscodeVideo("/home/wangshibiao/test/vr/vr1.mp4", "/home/wangshibiao/test/vr/vr1_out1.mp4")
	if err != nil {
		panic(err)
	}
}

/**
测试生成视频封面文件
*/
func TestGenVideoCover(t *testing.T) {
	err := util.GenVideoCover("/home/wangshibiao/test/vr/vr1.mp4", "/home/wangshibiao/test/vr/vr1.jpg")
	if err != nil {
		panic(err)
	}
}
