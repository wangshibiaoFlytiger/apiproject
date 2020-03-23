package cache

/**
初始化redis分布式缓存和本地缓存
*/
func Init() {
	//初始化redis
	initRedis()

	//初始化本地缓存
	initLocalCache()
}
