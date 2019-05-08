package test

import (
	"apiproject/cache"
	"apiproject/entity"
	m_video "apiproject/model/video"
	"fmt"
	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
	"testing"
	"time"
)

/**
测试redis
*/
func TestString(t *testing.T) {
	client := cache.RedisClient

	// 第三个参数是过期时间, 如果是0, 则表示没有过期时间.
	err := client.Set("name", "xys", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)

	// 这里设置过期时间.
	err = client.Set("age", "20", 1*time.Second).Err()
	if err != nil {
		panic(err)
	}

	client.Incr("age") // 自增
	client.Incr("age") // 自增
	client.Decr("age") // 自减

	val, err = client.Get("age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("age", val) // age 的值为21

	// 因为 key "age" 的过期时间是一秒钟, 因此当一秒后, 此 key 会自动被删除了.
	time.Sleep(1 * time.Second)
	val, err = client.Get("age").Result()
	if err != nil {
		// 因为 key "age" 已经过期了, 因此会有一个 cache: nil 的错误.
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println("age", val)
}

func TestSet(t *testing.T) {
	client := cache.RedisClient
	client.SAdd("blacklist", "Obama")     // 向 blacklist 中添加元素
	client.SAdd("blacklist", "Hillary")   // 再次添加
	client.SAdd("blacklist", "the Elder") // 添加新元素

	client.SAdd("whitelist", "the Elder") // 向 whitelist 添加元素

	// 判断元素是否在集合中
	isMember, err := client.SIsMember("blacklist", "Bush").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Is Bush in blacklist: ", isMember)

	// 求交集, 即既在黑名单中, 又在白名单中的元素
	names, err := client.SInter("blacklist", "whitelist").Result()
	if err != nil {
		panic(err)
	}
	// 获取到的元素是 "the Elder"
	fmt.Println("Inter result: ", names)

	// 获取指定集合的所有元素
	all, err := client.SMembers("blacklist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("All member: ", all)
}

func TestHash(t *testing.T) {
	client := cache.RedisClient

	client.HSet("user_xys", "name", "xys") // 向名称为 user_xys 的 hash 中添加元素 name
	client.HSet("user_xys", "age", "18")   // 向名称为 user_xys 的 hash 中添加元素 age

	// 批量地向名称为 user_test 的 hash 中添加元素 name 和 age
	client.HMSet("user_test", map[string]interface{}{"name": "test", "age": "20"})
	// 批量获取名为 user_test 的 hash 中的指定字段的值.
	fields, err := client.HMGet("user_test", "name", "age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("fields in user_test: ", fields)

	// 获取名为 user_xys 的 hash 中的字段个数
	length, err := client.HLen("user_xys").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("field count in user_xys: ", length) // 字段个数为2

	// 删除名为 user_test 的 age 字段
	client.HDel("user_test", "age")
	age, err := client.HGet("user_test", "age").Result()
	if err != nil {
		fmt.Printf("Get user_test age error: %v\n", err)
	} else {
		fmt.Println("user_test age is: ", age) // 字段个数为2test
	}
}

func TestZset(t *testing.T) {
	client := cache.RedisClient
	err := client.ZAdd("zset", redis.Z{Score: 1, Member: "one"}).Err()
	err = client.ZAdd("zset", redis.Z{Score: 2, Member: "two"}).Err()
	err = client.ZAdd("zset", redis.Z{Score: 3, Member: "three"}).Err()

	nowJsonTime := entity.JsonTime{time.Now()}
	video := m_video.Video{}
	video.ID = "id1"
	video.Title = "title1"
	video.CreatedAt = nowJsonTime

	//对象转为json后写入到redis
	videoJson, err := jsoniter.Marshal(video)
	fmt.Println(videoJson)
	err = client.ZAdd("zset", redis.Z{Score: 6, Member: videoJson}).Err()
	if err != nil {
		fmt.Print(err.Error())
	}

	zRange := client.ZRange("zset", 0, -1)
	for key, value := range zRange.Val() {
		fmt.Println(key, value)
	}
	fmt.Println(zRange)

	zRange = client.ZRange("zset", 2, 3)

	zRange = client.ZRange("zset", -2, -1)
}

func TestZsetJson(t *testing.T) {
	client := cache.RedisClient

	nowJsonTime := entity.JsonTime{time.Now()}
	video := m_video.Video{}
	video.ID = "id1"
	video.Title = "title1"
	video.CreatedAt = nowJsonTime

	//对象转为json后写入到redis
	videoJson, err := jsoniter.Marshal(video)
	fmt.Println(videoJson)
	err = client.ZAdd("zsetjson", redis.Z{Score: 6, Member: videoJson}).Err()
	if err != nil {
		fmt.Print(err.Error())
	}

	zRange := client.ZRange("zsetjson", 0, -1)
	for key, value := range zRange.Val() {
		fmt.Println(key, value)
		//redis读取的json转为对象
		videoTmp := m_video.Video{}
		jsoniter.UnmarshalFromString(value, &videoTmp)
		jsoniter.Unmarshal([]byte(value), &videoTmp)
		fmt.Println(videoTmp)
	}
}
