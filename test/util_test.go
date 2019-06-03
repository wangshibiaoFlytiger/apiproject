package test

import (
	"apiproject/util"
	"fmt"
	"testing"
)

/**
测试slice反转
*/
func TestSlice(t *testing.T) {
	slice := []interface{}{}
	slice = append(slice, 1)
	slice = append(slice, 3)
	slice = append(slice, 2)

	fmt.Println(slice)
	slice = util.ReverseSlice(slice)
	fmt.Println(slice)
}
