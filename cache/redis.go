package cache

import (
	"apiproject/config"
	"apiproject/log"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var RedisClient *redis.Client

// 创建 cache 客户端
func Init() {
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
		log.Logger.Error("redis客户端创建失败", zap.Error(err), zap.Any("pong", pong))
	}
	log.Logger.Info("RedisClient创建完成")
}
