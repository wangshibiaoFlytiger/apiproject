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
