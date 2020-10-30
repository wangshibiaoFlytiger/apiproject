package test

import (
	"apiproject/util"
	"github.com/patrickmn/go-cache"
	"testing"
	"time"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

/**
测试本地缓存库go-cache
*/
func TestGocache(t *testing.T) {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New(2*time.Second, 3*time.Second)

	c.Set("foo", "bar", cache.DefaultExpiration)
	c.Set("baz", 42, cache.NoExpiration)

	time.Sleep(1 * time.Second)

	if x, found := c.Get("foo"); found {
		println("foo:" + x.(string))
	}
	if x, found := c.Get("baz"); found {
		println("baz:" + util.IntToStr(x.(int)))
	}
}
