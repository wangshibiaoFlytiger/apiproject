package util

import (
	"math/rand"
	"time"
)

//生成范围在[start,end), 类型为int的随机数
func GenRandomInt(start int, end int) int {
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//生成随机数
	num := r.Intn((end - start)) + start

	return num
}

/**
随机获取1个数组元素
*/
func RandomElem(arr []interface{}) interface{} {
	randomIndex := GenRandomInt(0, len(arr))
	return arr[randomIndex]
}

//生成范围在[start,end), 类型为int的n个不重复随机数
func GenRandomIntList(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}

//生成范围在[start,end), 类型为int64的随机数
func GenRandomInt64(start int64, end int64) int64 {
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//生成随机数
	num := r.Int63n((end - start)) + start

	return num
}

//生成范围在[start,end), 类型为int64的n个不重复随机数
func GenRandomInt64List(start int64, end int64, count int) []int64 {
	//范围检查
	if end < start || (end-start) < IntToInt64(count) {
		return nil
	}

	//存放结果的slice
	nums := make([]int64, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Int63n((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}
