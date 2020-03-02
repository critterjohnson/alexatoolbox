package testgen

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRandomString(t *testing.T) {
	Convey("When RandomString is called", t, func() {
		str := RandomString()

		Convey("it should return a string of length 16", func() {
			So(len(str), ShouldEqual, 16)
		})
	})
}

func TestRandomStruct(t *testing.T) {
	Convey("When RandomStruct is called", t, func() {
		type RandStruct struct {
			Value1 string
			Value2 string
		}

		structure := RandomStruct(reflect.TypeOf(RandStruct{}), t).(RandStruct)

		Convey("it should return a populated struct", func() {
			So(structure, ShouldNotBeEmpty)
		})
	})
}

func TestRandomRequest(t *testing.T) {
	Convey("When RandomRequest is called", t, func() {
		req := RandomRequest(t)

		Convey("It should return a populated Request object", func() {
			So(req, ShouldNotBeEmpty)
		})
	})
}

func TestRandomResponse(t *testing.T) {
	Convey("When RandomResponse is called", t, func() {
		res := RandomResponse(t)

		Convey("It should return a populated Response object", func() {
			So(res, ShouldNotBeEmpty)
		})
	})
}
