package cache

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

/**
初始化redis分布式缓存和本地缓存
*/
func Init() {
	//初始化redis
	initRedis()

	//初始化本地缓存
	initLocalCache()
}
