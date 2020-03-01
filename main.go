package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/critterjohnson/alexatoolbox/handlers"
	"github.com/critterjohnson/alexatoolbox/request"
	"github.com/critterjohnson/alexatoolbox/response"
)

func SayHello(request request.Request) (response.Response, error) {
	return response.NewBuilder().
		WithTextOutputSpeech("hello, " + request.RequestBody.Intent.Slots["name"].Value).
		Build(), nil
}

func LambdaHandler(ctx context.Context, request request.Request) (response.Response, error) {
	requestHandler := handlers.NewRequestHandler()
	requestHandler.AddIntentRequestHandler("SayHello", SayHello)
	response := requestHandler.Handle(request)
	return response, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
