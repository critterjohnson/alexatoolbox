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
