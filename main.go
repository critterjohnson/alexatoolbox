package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/critterjohnson/go-ask/alexarequest"
	"github.com/critterjohnson/go-ask/alexaresponse"
	"github.com/critterjohnson/go-ask/handlers"
)

func SayHello(request alexarequest.Request) (alexaresponse.Response, error) {
	return alexaresponse.NewBuilder().
		WithTextOutputSpeech("hello, " + request.RequestBody.Intent.Slots["name"].Value).
		Build(), nil
}

func LambdaHandler(ctx context.Context, request alexarequest.Request) (alexaresponse.Response, error) {
	requestHandler := handlers.NewRequestHandler()
	requestHandler.AddIntentRequestHandler("SayHello", SayHello)
	response, err := requestHandler.Handle(request)
	return response, err
}

func main() {
	data, err := ioutil.ReadFile("sonnet138.json")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	request := alexarequest.Request{}
	if err = json.Unmarshal(data, &request); err != nil {
		fmt.Println("Error unmarshalling JSON", err)
	}
	fmt.Printf("%+v", request.Session)
}
