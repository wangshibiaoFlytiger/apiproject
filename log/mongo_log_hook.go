package log

import (
	"apiproject/config"
	"apiproject/mongo"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

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
	err = collection.Insert(object)
	if err != nil {
		HookLogger.Error("插入日志到mongo, 插入mongo异常", zap.Any("data", data), zap.Error(err))
		return
	}

	HookLogger.Info("插入日志到mongo, 完成")
	return
}
