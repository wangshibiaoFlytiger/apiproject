package cache

import (
	"apiproject/config"
	"apiproject/log"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var RedisClient *redis.Client

// 创建 cache 客户端
func init() {
	redisHost, err := config.Config.GetValue("redis", "host")
	if err != nil {
		log.Logger.Error("读取redis host配置出错", zap.Error(err))
	}
	redisPort, err := config.Config.GetValue("redis", "port")
	if err != nil {
		log.Logger.Error("读取redis port配置出错", zap.Error(err))
	}
	redisPassword, err := config.Config.GetValue("redis", "password")
	if err != nil {
		log.Logger.Error("读取redis redisPassword配置出错", zap.Error(err))
	}
	redisDb, err := config.Config.Int("redis", "db")
	if err != nil {
		log.Logger.Error("读取redis redisDb配置出错", zap.Error(err))
	}
	poolSize, err := config.Config.Int("redis", "pool.size")
	if err != nil {
		log.Logger.Error("读取redis poolSize配置出错", zap.Error(err))
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       redisDb,
		PoolSize: poolSize,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 cache 服务器
	pong, err := RedisClient.Ping().Result()
	if err != nil {
		log.Logger.Error("redis客户端创建失败", zap.Error(err), zap.Any("pong", pong))
	}
	log.Logger.Info("RedisClient创建完成")
}
