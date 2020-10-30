package util

import (
	"github.com/vansante/go-ffprobe"
	"log"
	"time"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

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

/**
转码视频
*/
func TranscodeVideo(inputPath string, outputPath string) error {
	stdOut, stdErr, err := ExecCmd("ffmpeg", "-y", "-i", inputPath, "-c:v", "libx264", "-c:a", "aac", outputPath)
	if err != nil {
		log.Panicln("转码视频异常", err, inputPath, outputPath)
		return err
	}

	log.Println("转码视频完成", inputPath, outputPath, stdOut, stdErr)
	return nil
}

/**
生成视频封面文件
*/
func GenVideoCover(videoPath string, coverPath string) error {
	_, _, seconds := GetVideoInfo(videoPath)
	stdOut, stdErr, err := ExecCmd("ffmpeg", "-y", "-ss", IntToStr(Float64ToInt(seconds)/2), "-i", videoPath, "-r", "1", "-vframes", "1", "-an", "-vcodec", "mjpeg", coverPath)
	if err != nil {
		log.Panicln("生成视频封面文件异常", err, videoPath, coverPath)
		return err
	}

	log.Println("生成视频封面文件完成", videoPath, coverPath, stdOut, stdErr)
	return nil
}
