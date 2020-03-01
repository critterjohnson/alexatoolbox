package testgen

import (
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"testing/quick"
	"time"

	"github.com/critterjohnson/go-ask/request"
	"github.com/critterjohnson/go-ask/response"
)

// RandomString generates a new random string of length 16.
func RandomString() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune(
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
			"abcdefghijklmnopqrstuvwxyz" +
			"0123456789")
	length := 16
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}

// RandomStruct returns a randomly generated request.Request object.
func RandomStruct(typ reflect.Type, t *testing.T) interface{} {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	obj, ok := quick.Value(typ, random)
	if !ok {
		t.Errorf("failed to generate struct of type %s", typ)
	}
	objInterface := obj.Interface()
	return objInterface
}

// RandomRequest returns a randomly populated request object, without panicking.
func RandomRequest(t *testing.T) request.Request {
	req := request.Request{
		Version: RandomString(),
		Session: request.Session{
			ID:          RandomString(),
			Application: RandomStruct(reflect.TypeOf(request.Application{}), t).(request.Application),
			User:        RandomStruct(reflect.TypeOf(request.User{}), t).(request.User),
		},
		Context: request.Context{
			Application:    RandomStruct(reflect.TypeOf(request.Application{}), t).(request.Application),
			User:           RandomStruct(reflect.TypeOf(request.User{}), t).(request.User),
			Person:         RandomStruct(reflect.TypeOf(request.Person{}), t).(request.Person),
			APIEndpoint:    RandomString(),
			APIAccessToken: RandomString(),
			AudioPlayer:    RandomStruct(reflect.TypeOf(request.AudioPlayer{}), t).(request.AudioPlayer),
		},
		RequestBody: RandomStruct(reflect.TypeOf(request.RequestBody{}), t).(request.RequestBody),
	}
	return req
}

// RandomResponse returns a randomly populated resonse object, without panicking.
func RandomResponse(t *testing.T) response.Response {
	outputSpeech := RandomStruct(reflect.TypeOf(response.OutputSpeech{}), t).(response.OutputSpeech)
	card := RandomStruct(reflect.TypeOf(response.Card{}), t).(response.Card)
	reprompt := RandomStruct(reflect.TypeOf(response.Reprompt{}), t).(response.Reprompt)

	res := response.Response{
		Version: RandomString(),
		Response: &response.ResponseBody{
			OutputSpeech: &outputSpeech,
			Card:         &card,
			Reprompt:     &reprompt,
		},
	}
	return res
}
