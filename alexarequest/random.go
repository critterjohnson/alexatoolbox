package alexarequest

import (
	"reflect"
	"testing"

	"github.com/critterjohnson/go-ask/testgen"
)

// RandomRequest returns a randomly populated request object, without panicking.
func RandomRequest(t *testing.T) Request {
	req := Request{
		Version: testgen.RandomString(),
		Session: Session{
			ID:          testgen.RandomString(),
			Application: testgen.RandomStruct(reflect.TypeOf(Application{}), t).(Application),
			User:        testgen.RandomStruct(reflect.TypeOf(User{}), t).(User),
		},
		Context: Context{
			Application:    testgen.RandomStruct(reflect.TypeOf(Application{}), t).(Application),
			User:           testgen.RandomStruct(reflect.TypeOf(User{}), t).(User),
			Person:         testgen.RandomStruct(reflect.TypeOf(Person{}), t).(Person),
			APIEndpoint:    testgen.RandomString(),
			APIAccessToken: testgen.RandomString(),
			AudioPlayer:    testgen.RandomStruct(reflect.TypeOf(AudioPlayer{}), t).(AudioPlayer),
		},
		RequestBody: testgen.RandomStruct(reflect.TypeOf(Body{}), t).(Body),
	}
	return req
}
