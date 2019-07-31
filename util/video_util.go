package util

import (
	"github.com/vansante/go-ffprobe"
	"log"
	"time"
)

/**
获取视频信息
*/
func GetVideoInfo(videoPath string) (width int, height int, seconds float64) {
	data, err := ffprobe.GetProbeData(videoPath, 5000*time.Millisecond)
	if err != nil {
		log.Panicf("获取视频信息异常: %v", err)
	}

	firstVideoStream := data.GetFirstVideoStream()
	return firstVideoStream.Width, firstVideoStream.Height, data.Format.DurationSeconds
}
