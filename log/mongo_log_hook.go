package log

import (
	"apiproject/config"
	"apiproject/mongo"
	"github.com/globalsign/mgo"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"time"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

type MongoLogHook struct {
}

/**
实现MongoLogHook的io.writer接口方法, 使当前对象可以作为zap的hook使用
*/
func (this *MongoLogHook) Write(data []byte) (n int, err error) {
	if err = this.InsertLogToMongo(data); err != nil {
		return 0, err
	}

	return
}

/**
插入日志到mongo
*/
func (this *MongoLogHook) InsertLogToMongo(data []byte) (err error) {
	sessionClone := mongo.GetMongoSessionClone()
	defer sessionClone.Close()
	db := sessionClone.DB(config.GlobalConfig.MongoDatabase)

	collection := db.C(config.GlobalConfig.LogMongoCollection)
	var object interface{}
	if err = jsoniter.Unmarshal(data, &object); err != nil {
		HookLogger.Error("插入日志到mongo, json解析异常", zap.Any("data", data), zap.Error(err))
		return
	}

	//转为map类型
	dataMap := object.(map[string]interface{})

	//mongo日志过期, 则自动清理
	dateTimeFieldName := "dateTime"
	dataMap[dateTimeFieldName] = time.Now()
	datetimeIndexUsedForExpire := mgo.Index{
		Key:         []string{dateTimeFieldName},
		ExpireAfter: time.Duration(config.GlobalConfig.LogMongoExpireAfterSeconds) * time.Second}
	if err = collection.EnsureIndex(datetimeIndexUsedForExpire); err != nil {
		HookLogger.Error("插入日志到mongo, 创建时间索引异常", zap.Any("data", dataMap), zap.Error(err))
		return
	}

	err = collection.Insert(dataMap)
	if err != nil {
		HookLogger.Error("插入日志到mongo, 插入mongo异常", zap.Any("data", dataMap), zap.Error(err))
		return
	}

	HookLogger.Info("插入日志到mongo, 完成")
	return
}
