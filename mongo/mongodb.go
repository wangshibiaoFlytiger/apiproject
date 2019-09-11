package mongo

import (
	"apiproject/config"
	"github.com/globalsign/mgo"
	"log"
	"time"
)

var globalMongoSession *mgo.Session

/**
初始化mongo
*/
func Init() {
	if !config.GlobalConfig.MongoSwitch {
		return
	}

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
		log.Panicln("初始化mongo, 异常", err)
	}

	log.Println("初始化mongo, 完成")
}

/**
获取session的clone
*/
func GetMongoSessionClone() *mgo.Session {
	return globalMongoSession.Clone()
}
