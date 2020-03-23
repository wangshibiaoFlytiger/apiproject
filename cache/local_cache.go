package cache

import (
	"apiproject/config"
	"github.com/patrickmn/go-cache"
	"time"
)

//本地全局缓存
var LocalCache *cache.Cache

/**
初始化本地缓存
*/
func initLocalCache() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	LocalCache = cache.New(time.Duration(config.GlobalConfig.LocalcacheDefaultExpirationSeconds)*time.Second, time.Duration(config.GlobalConfig.LocalcacheCleanupIntervalSeconds)*time.Second)
}
