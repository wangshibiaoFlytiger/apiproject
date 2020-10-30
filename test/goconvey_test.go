package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	convey.Convey("Given some integer with a starting value", t, func() {
		x := 1

		convey.Convey("When the integer is incremented", func() {
			x++

			convey.Convey("The value should be greater by one", func() {
				convey.So(x, convey.ShouldEqual, 2)
			})
		})
	})
}
