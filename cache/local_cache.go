package cache

import (
	"apiproject/config"
	"github.com/patrickmn/go-cache"
	"time"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

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
