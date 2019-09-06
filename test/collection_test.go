package test

import (
	"apiproject/util"
	"fmt"
	"testing"
)

/**
测试集合的包含计算
*/
func TestCollectionContain(t *testing.T) {
	//数组
	a := 1
	b := [3]int{1, 2, 3}
	fmt.Println(util.Contains(b, a))

	//slice
	c := 1
	d := []int{1, 2, 3}
	fmt.Println(util.Contains(d, c))

	//map
	var e = map[int]string{1: "1", 2: "2"}
	fmt.Println(util.Contains(e, 2))
}
