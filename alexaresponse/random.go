package alexaresponse

import (
	"reflect"
	"testing"

	"github.com/critterjohnson/go-ask/testgen"
)

// RandomResponse returns a randomly populated resonse object, without panicking.
func RandomResponse(t *testing.T) Response {
	outputSpeech := testgen.RandomStruct(reflect.TypeOf(OutputSpeech{}), t).(OutputSpeech)
	card := testgen.RandomStruct(reflect.TypeOf(Card{}), t).(Card)
	reprompt := testgen.RandomStruct(reflect.TypeOf(Reprompt{}), t).(Reprompt)

	res := Response{
		Version: testgen.RandomString(),
		Response: &Body{
			OutputSpeech: &outputSpeech,
			Card:         &card,
			Reprompt:     &reprompt,
		},
	}
	return res
}
