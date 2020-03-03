package alexarequest

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRandomRequest(t *testing.T) {
	Convey("When RandomRequest is called", t, func() {
		req := RandomRequest(t)

		Convey("It should return a populated Request object", func() {
			So(req, ShouldNotBeEmpty)
		})
	})
}
