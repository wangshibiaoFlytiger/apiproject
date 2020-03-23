package cache

import (
	"apiproject/config"
	"github.com/go-redis/redis"
	"log"
	"strconv"
	"time"
)

var RedisClient *redis.Client

/**
初始化redis
*/
func initRedis() {
	if !config.GlobalConfig.RedisSwitch {
		return
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:         config.GlobalConfig.RedisHost + ":" + strconv.Itoa(config.GlobalConfig.RedisPort),
		Password:     config.GlobalConfig.RedisPassword,
		DB:           config.GlobalConfig.RedisDb,
		PoolSize:     config.GlobalConfig.RedisPoolSize,
		MaxRetries:   2,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		MinIdleConns: 50,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 cache 服务器
	pong, err := RedisClient.Ping().Result()
	if err != nil {
		log.Panicln("初始化redis, 异常", pong, err)
	}

	log.Println("初始化redis, 完成")
}
