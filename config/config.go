package config

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/go-ini/ini"
)

/**
系统配置
*/
type Config struct {
	//运行环境:dev,test,pro
	Profile string

	ServicePort               int  `ini:"service.port"`
	ServiceApiResponseEncrypt bool `ini:"service.api.response.encrypt"`

	MysqlSwitch       bool   `ini:"mysql.switch"`
	MysqlUrl          string `ini:"mysql.url"`
	MysqlMaxIdleCount int    `ini:"mysql.max.idle.count"`
	MysqlMaxOpenCount int    `ini:"mysql.max.open.count"`

	RedisSwitch   bool   `ini:"redis.switch"`
	RedisHost     string `ini:"redis.host"`
	RedisPort     int    `ini:"redis.port"`
	RedisPassword string `ini:"redis.password"`
	RedisDb       int    `ini:"redis.db"`
	RedisPoolSize int    `ini:"redis.pool.size"`

	MongoSwitch    bool   `ini:"mongo.switch"`
	MongoAddr      string `ini:"mongo.addr"`
	MongoDatabase  string `ini:"mongo.database"`
	MongoUserName  string `ini:"mongo.username"`
	MongoPassword  string `ini:"mongo.password"`
	MongoPoolLimit int    `ini:"mongo.pool.limit"`

	LogDir             string `ini:"log.dir"`
	LogMaxDayCount     int    `ini:"log.max.day.count"`
	LogKafkaHookSwitch bool   `ini:"log.kafka.hook.switch"`
	LogKafkaTopic      string `ini:"log.kafka.topic"`
	LogMongoHookSwitch bool   `ini:"log.mongo.hook.switch"`
	LogMongoCollection string `ini:"log.mongo.collection"`

	KafkaSwitch bool   `ini:"kafka.switch"`
	KafkaBroker string `ini:"kafka.broker"`

	WxpayH5Appid     string `ini:"wxpay.h5.appid"`
	WxpayH5Mchid     string `ini:"wxpay.h5.mchid"`
	WxpayH5Apikey    string `ini:"wxpay.h5.apikey"`
	WxpayH5Notifyurl string `ini:"wxpay.h5.notifyurl"`

	IplocationQqwryPath string `ini:"iplocation.qqwry.path"`
}

var GlobalConfig Config = Config{}

func Init() {
	//通过go.rice读取配置文件的内容
	box := rice.MustFindBox("./config_file")
	configFileSubPath := "config_" + GlobalConfig.Profile + ".ini"
	configContent := box.MustBytes(configFileSubPath)

	//加载配置文件
	cfg, err := ini.Load(configContent)

	err = cfg.Section("service").MapTo(&GlobalConfig)
	if err != nil {
		panic(err)
	}

	err = cfg.Section("mysql").MapTo(&GlobalConfig)
	if err != nil {
		panic(err)
	}

	err = cfg.Section("log").MapTo(&GlobalConfig)
	if err != nil {
		panic(err)
	}

	err = cfg.Section("redis").MapTo(&GlobalConfig)
	if err != nil {
		panic(err)
	}

	err = cfg.Section("mongo").MapTo(&GlobalConfig)
	if err != nil {
		panic(err)
	}

	err = cfg.Section("kafka").MapTo(&GlobalConfig)
	if err != nil {
		panic(err)
	}

	err = cfg.Section("wxpay").MapTo(&GlobalConfig)
	if err != nil {
		panic(err)
	}

	err = cfg.Section("iplocation").MapTo(&GlobalConfig)
	if err != nil {
		panic(err)
	}
}
