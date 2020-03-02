package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/critterjohnson/go-ask/handlers"
	"github.com/critterjohnson/go-ask/request"
	"github.com/critterjohnson/go-ask/response"
)

func SayHello(request request.Request) (response.Response, error) {
	return response.NewBuilder().
		WithTextOutputSpeech("hello, " + request.RequestBody.Intent.Slots["name"].Value).
		Build(), nil
}

func LambdaHandler(ctx context.Context, request request.Request) (response.Response, error) {
	requestHandler := handlers.NewRequestHandler()
	requestHandler.AddIntentRequestHandler("SayHello", SayHello)
	response, err := requestHandler.Handle(request)
	return response, err
}

func main() {
	lambda.Start(LambdaHandler)
}
