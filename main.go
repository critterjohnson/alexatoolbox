package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
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
	lambda.Start(LambdaHandler)
}
