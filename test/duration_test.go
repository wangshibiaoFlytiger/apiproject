package test

import (
	"apiproject/util"
	"testing"
	"time"
)

/**
测试duration格式化
*/
func TestDurationHuman(t *testing.T) {
	duration := (25 * time.Hour) + (22 * time.Minute) + (63 * time.Second)
	util.GetDurationHuman(duration)
}
