package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/critterjohnson/go-ask/alexarequest"
	"github.com/critterjohnson/go-ask/alexaresponse"
	"github.com/critterjohnson/go-ask/handlers"
)

// Launch is the handler for LaunchRequests.
func Launch(request alexarequest.Request) (alexaresponse.Response, error) {
	return alexaresponse.NewBuilder().
		WithTextOutputSpeech("Welcome to GoHello! Please say hello and introduce yourself with your name.").
		ShouldEndSession(false).
		Build(), nil
}

// SayHello is the handler for a SayHello Intent.
func SayHello(request alexarequest.Request) (alexaresponse.Response, error) {
	return alexaresponse.NewBuilder().
		WithTextOutputSpeech("Hello, " + request.RequestBody.Intent.Slots["name"].Value).
		Build(), nil
}

// End is the handler for SessionEndedRequests. noop.
func End(request alexarequest.Request) error {
	return nil
}

// LambdaHandler is the AWS Lambda event handler for incoming Alexa requests.
func LambdaHandler(ctx context.Context, request alexarequest.Request) (alexaresponse.Response, error) {
	requestHandler := handlers.NewRequestHandler()
	requestHandler.AddLaunchRequestHandler(Launch)
	requestHandler.AddIntentRequestHandler("SayHello", SayHello)
	requestHandler.AddSessionEndedRequestHandler(End)
	response, err := requestHandler.Handle(request)
	return response, err
}

func main() {
	lambda.Start(LambdaHandler)
}
