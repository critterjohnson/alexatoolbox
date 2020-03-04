package alexarequest

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/critterjohnson/go-ask/testgen"
	. "github.com/smartystreets/goconvey/convey"
)

func trim(str string) string {
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, " ", "", -1)
	return str
}

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

		Convey("it should return the attributes JSON", func() {
			trimmedData := trim(string(data))
			trimmedAttrs := trim(string(attributes))

			So(trimmedData, ShouldEqual, trimmedAttrs)
		})
	})
}

func TestUnmarshalJSON(t *testing.T) {
	var session Session

	model := testgen.RandomString()
	color := testgen.RandomString()
	horsepower := testgen.RandomString()
	attributes := []byte(fmt.Sprintf(`
		{
			"model": "%s",
			"color": "%s",
			"horsepower": "%s"
		}
	`, model, color, horsepower))

	sessionID := testgen.RandomString()
	applicationID := testgen.RandomString()
	userID := testgen.RandomString()
	accessToken := testgen.RandomString()
	perm := testgen.RandomString()

	jsonData := []byte(fmt.Sprintf(`
	{
		"new": true,
		"sessionId": "%s",
		"application": {
			"applicationId": "%s"
		},
		"user": {
			"userId": "%s",
			"accessToken": "%s",
			"permissions": {
				"perm": "%s"
			}
		},
		"attributes": %s
	}
	`, sessionID, applicationID, userID, accessToken, perm, string(attributes)))

	Convey("When UnmarshalJSON is called", t, func() {
		err := json.Unmarshal(jsonData, &session)

		Convey("it should return a nil error", func() {
			So(err, ShouldBeNil)
		})
		Convey("it should take on the json data", func() {
			So(session.ID, ShouldEqual, sessionID)
			So(session.Application.ID, ShouldEqual, applicationID)
			So(session.User.ID, ShouldEqual, userID)
			So(session.User.AccessToken, ShouldEqual, accessToken)
			So(session.User.Permissions["perm"], ShouldEqual, perm)
			So(trim(string(session.attributeData)), ShouldEqual, trim(string(attributes)))
		})
	})
}

func TestUnmarshalAttributes(t *testing.T) {
	var session Session

	var attributes struct {
		Model      string `json:"model"`
		Color      string `json:"color"`
		Horsepower string `json:"horsepower"`
	}

	model := testgen.RandomString()
	color := testgen.RandomString()
	horsepower := testgen.RandomString()

	jsonData := []byte(fmt.Sprintf(`
	{
		"new": true,
		"sessionId": "id",
		"application": {
			"applicationId": "id"
		},
		"user": {
			"userId": "id",
			"accessToken": "token",
			"permissions": {
				"perm": "perm"
			}
		},
		"attributes": {
			"model": "%s",
			"color": "%s",
			"horsepower": "%s"
		}
	}
	`, model, color, horsepower))

	if err := json.Unmarshal(jsonData, &session); err != nil {
		t.Error("Failed to unmarshal Session.")
	}

	Convey("When UnmarshalAttributes is called", t, func() {
		err := session.UnmarshalAttributes(&attributes)

		Convey("it should return a nil error", func() {
			So(err, ShouldBeNil)
		})

		Convey("it should return the Attributes object", func() {
			So(attributes, ShouldNotBeEmpty)

			Convey("with the proper data", func() {
				So(attributes.Model, ShouldEqual, model)
				So(attributes.Color, ShouldEqual, color)
				So(attributes.Horsepower, ShouldEqual, horsepower)
			})
		})
	})
}
