package bootstrap

import (
	"apiproject/cache"
	"apiproject/config"
	c_video "apiproject/controller/video"
	"apiproject/dao"
	"apiproject/log"
	s_video "apiproject/service/video"
)

/**
系统初始化
*/
func Init() {
	config.Init()
	log.Init()

	cache.Init()

	//dao层初始化
	dao.Init()

	//service层初始化
	s_video.Init()

	//controller层初始化
	c_video.Init()
}
