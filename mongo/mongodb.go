package mongo

import (
	"apiproject/config"
	"apiproject/log"
	"github.com/globalsign/mgo"
	"go.uber.org/zap"
	"time"
)

var globalMongoSession *mgo.Session

func Init() {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{config.GlobalConfig.MongoAddr},
		Direct:    false,
		Timeout:   time.Second * 1,
		Database:  config.GlobalConfig.MongoDatabase,
		Source:    "admin",
		Username:  config.GlobalConfig.MongoUserName,
		Password:  config.GlobalConfig.MongoPassword,
		PoolLimit: config.GlobalConfig.MongoPoolLimit,
	}

	var err error
	globalMongoSession, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Logger.Error("mongo初始化异常", zap.Error(err))
	}
}

/**
获取session的clone
*/
func GetMongoSessionClone() *mgo.Session {
	return globalMongoSession.Clone()
}
