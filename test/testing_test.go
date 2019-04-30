package test

import (
	"apiproject/log"
	s_video "apiproject/service/video"
	"github.com/Unknwon/goconfig"
	"go.uber.org/zap"
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

/**
测试配置文件解析
*/
func TestParseConfigFile(t *testing.T) {
	c, err := goconfig.LoadConfigFile("config/config.ini")
	if err != nil {
		log.Logger.Error("测试异常", zap.Error(err))
		return
	}

	mysqlUrl, err := c.GetValue("mysql", "url")
	if err != nil {
		log.Logger.Error("测试异常", zap.Error(err))
		return
	}

	log.Logger.Info("解析配置文件", zap.String("mysqlUrl", mysqlUrl))
}
