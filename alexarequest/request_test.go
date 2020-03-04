package alexarequest

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExtractAttributes(t *testing.T) {
	Convey("When TestExtractAttributes is called", t, func() {
		jsonData := []byte(`{
			"fakeKey": "fakeVal",
			"fakeKey2": "fakeVal2",
			"attributes": {
				"obj1": {
					"key": "val"
				},
				"key": "val"
			}	
		}`)
		attributes := []byte(`{
			"obj1": {
				"key": "val"
			},
			"key": "val"
		}`)

		data := extractAttributes(jsonData)
		fmt.Println(string(attributes))

		Convey("it should return the attributes JSON", func() {
			trim := func(str string) string {
				str = strings.Replace(str, "\n", "", -1)
				str = strings.Replace(str, "\t", "", -1)
				str = strings.Replace(str, " ", "", -1)
				return str
			}
			trimmedData := trim(string(data))
			trimmedAttrs := trim(string(attributes))
			fmt.Println(trimmedAttrs)

			So(trimmedData, ShouldEqual, trimmedAttrs)
		})
	})
}
