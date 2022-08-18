package test

import (
	"apiproject/util"
	"fmt"
	"testing"
	"time"
	. "github.com/smartystreets/goconvey/convey"
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

func TestSpec(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}