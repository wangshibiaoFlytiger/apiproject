package test

import (
	"apiproject/util"
	"fmt"
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
测试duration格式化
*/
func TestDurationHuman(t *testing.T) {
	duration := (25 * time.Hour) + (22 * time.Minute) + (63 * time.Second)
	durationHuman := util.GetDurationHuman(duration)
	fmt.Println(durationHuman)
}
