package alexaresponse

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRandomResponse(t *testing.T) {
	Convey("When RandomResponse is called", t, func() {
		res := RandomResponse(t)

		Convey("It should return a populated Response object", func() {
			So(res, ShouldNotBeEmpty)
		})
	})
}
