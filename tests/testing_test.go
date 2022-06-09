package testing

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// * testing
func TestSomething(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented by 2", func() {
			x += 2

			Convey("The value should be equal to 3!", func() {
				So(x, ShouldEqual, 3)
			})
		})
	})
}
