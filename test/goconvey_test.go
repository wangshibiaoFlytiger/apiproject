package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

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
