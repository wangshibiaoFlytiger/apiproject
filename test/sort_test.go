package test

import (
	"fmt"
	"sort"
	"testing"
)

/**
测试slice自定义排序
*/
func TestSliceSort(t *testing.T) {
	a := []int{1, 3, 2}
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	fmt.Println(a)

	b := []bool{true, false, false}
	sort.Slice(b, func(i, j int) bool {
		if b[i] == true && b[j] == false {
			return true
		}
		return false
	})
	fmt.Println(b)
}
