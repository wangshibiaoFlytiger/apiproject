package test

import (
	s_video "apiproject/service/video"
	"fmt"
	"github.com/Unknwon/goconfig"
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
	c, err := goconfig.LoadConfigFile("../config/config.ini")
	mysqlUrl, err := c.GetValue("mysql", "url")
	fmt.Println(mysqlUrl, err)
}
