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
	a := []string{"1", "3", "2"}
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	fmt.Println(a)
}
