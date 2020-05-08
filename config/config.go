package config

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/go-ini/ini"
)

/**
系统配置
*/
type Config struct {
	//运行环境:back-dev,back-test,front-dev,front-test
	Profile string `json:"profile"`

	ServiceName                                  string `ini:"service.name" json:"serviceName"`
	ServicePort                                  int    `ini:"service.port" json:"servicePort"`
	ServiceApiResponseEncrypt                    bool   `ini:"service.api.response.encrypt" json:"serviceApiResponseEncrypt"`
	ServiceRequestbodyLimitByteCount             int64  `ini:"service.requestbody.limit.byte.count" json:"serviceRequestbodyLimitByteCount"`
	ServiceResponsecacheDefaultExpirationSeconds int    `ini:"service.responseCache.defaultExpiration.seconds" json:"serviceResponsecacheDefaultExpirationSeconds"`

	MysqlSwitch       bool   `ini:"mysql.switch" json:"mysqlSwitch"`
	MysqlUrl          string `ini:"mysql.url" json:"mysqlUrl"`
	MysqlMaxIdleCount int    `ini:"mysql.max.idle.count" json:"mysqlMaxIdleCount"`
	MysqlMaxOpenCount int    `ini:"mysql.max.open.count" json:"mysqlMaxOpenCount"`

	RedisSwitch   bool   `ini:"redis.switch" json:"redisSwitch"`
	RedisHost     string `ini:"redis.host" json:"redisHost"`
	RedisPort     int    `ini:"redis.port" json:"redisPort"`
	RedisPassword string `ini:"redis.password" json:"redisPassword"`
	RedisDb       int    `ini:"redis.db" json:"redisDb"`
	RedisPoolSize int    `ini:"redis.pool.size" json:"redisPoolSize"`

	LocalcacheDefaultExpirationSeconds int `ini:"localcache.defaultExpiration.seconds" json:"localcacheDefaultExpirationSeconds"`
	LocalcacheCleanupIntervalSeconds   int `ini:"localcache.cleanupInterval.seconds" json:"localcacheCleanupIntervalSeconds"`

	MongoSwitch    bool   `ini:"mongo.switch" json:"mongoSwitch"`
	MongoAddr      string `ini:"mongo.addr" json:"mongoAddr"`
	MongoDatabase  string `ini:"mongo.database" json:"mongoDatabase"`
	MongoUserName  string `ini:"mongo.username" json:"mongoUserName"`
	MongoPassword  string `ini:"mongo.password" json:"mongoPassword"`
	MongoPoolLimit int    `ini:"mongo.pool.limit" json:"mongoPoolLimit"`

	LogDir                     string `ini:"log.dir" json:"logDir"`
	LogMaxDayCount             int    `ini:"log.max.day.count" json:"logMaxDayCount"`
	LogKafkaHookSwitch         bool   `ini:"log.kafka.hook.switch" json:"logKafkaHookSwitch"`
	LogKafkaTopic              string `ini:"log.kafka.topic" json:"logKafkaTopic"`
	LogMongoHookSwitch         bool   `ini:"log.mongo.hook.switch" json:"logMongoHookSwitch"`
	LogMongoCollection         string `ini:"log.mongo.collection" json:"logMongoCollection"`
	LogMongoExpireAfterSeconds int    `ini:"log.mongo.expire.after.seconds" json:"logMongoExpireAfterSeconds"`

	KafkaSwitch bool   `ini:"kafka.switch" json:"kafkaSwitch"`
	KafkaBroker string `ini:"kafka.broker" json:"kafkaBroker"`

	WxpayH5Appid     string `ini:"wxpay.h5.appid" json:"wxpayH5Appid"`
	WxpayH5Mchid     string `ini:"wxpay.h5.mchid" json:"wxpayH5Mchid"`
	WxpayH5Apikey    string `ini:"wxpay.h5.apikey" json:"wxpayH5Apikey"`
	WxpayH5Notifyurl string `ini:"wxpay.h5.notifyurl" json:"wxpayH5Notifyurl"`

	IplocationQqwryPath string `ini:"iplocation.qqwry.path" json:"iplocationQqwryPath"`

	//反向代理列表
	ReverseproxyList string `ini:"reverseproxy.list" json:"reverseproxyList"`
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

	err = cfg.Section("localcache").MapTo(&GlobalConfig)
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

	err = cfg.Section("reverseproxy").MapTo(&GlobalConfig)
	if err != nil {
		panic(err)
	}
}
