package testgen

import (
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"testing/quick"
	"time"
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
