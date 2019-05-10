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

	MysqlUrl          string `ini:"mysql.url"`
	MysqlMaxIdleCount int    `ini:"mysql.max.idle.count"`
	MysqlMaxOpenCount int    `ini:"mysql.max.open.count"`

	RedisHost     string `ini:"redis.host"`
	RedisPort     int    `ini:"redis.port"`
	RedisPassword string `ini:"redis.password"`
	RedisDb       int    `ini:"redis.db"`
	RedisPoolSize int    `ini:"redis.pool.size"`

	LogPath  string `ini:"log.path"`
	Loglevel string `ini:"log.level"`
}

var GlobalConfig Config = Config{}

func Init() {
	//通过go.rice读取配置文件的内容
	box := rice.MustFindBox("./")
	configFileSubPath := "config_" + GlobalConfig.Profile + ".ini"
	configContent := box.MustBytes(configFileSubPath)

	//加载配置文件
	cfg, err := ini.Load(configContent)

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
}
