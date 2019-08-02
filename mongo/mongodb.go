package mongo

import (
	"apiproject/config"
	"apiproject/log"
	"github.com/globalsign/mgo"
	"go.uber.org/zap"
)

var MyMongodb *mgo.Database

func Init() {
	url := config.GlobalConfig.MongoUrl
	// 解析MongoDB参数
	dialInfo, err := mgo.ParseURL(url)
	// 1、连接MongoDB
	session, err := mgo.Dial(url)
	if err != nil {
		log.Logger.Error("mongo初始化异常", zap.Error(err))
	}

	// 2、选择数据库
	MyMongodb = session.DB(dialInfo.Database)
}
