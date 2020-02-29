package handlers

import (
	"github.com/critterjohnson/alexatoolbox/request"
	"github.com/critterjohnson/alexatoolbox/response"
)

// RequestHandler handles incoming requests and sends them to user-defined handlers.
type RequestHandler struct {
	launchRequestHandler       func(request.Request) (response.Response, error)
	intentRequestHandlers      map[string]func(request.Request) (response.Response, error)
	sessionEndedRequestHandler func(request.Request) error
	errorHandler               func(request.Request, error) response.Response
}

// NewRequestHandler returns a new request handler.
func NewRequestHandler() *RequestHandler {
	return &RequestHandler{
		intentRequestHandlers: make(map[string]func(request.Request) (response.Response, error)),
	}
}

// Handle handles an incoming request by calling the user-defined handlers.
func (r *RequestHandler) Handle(request request.Request) response.Response {
	var response response.Response
	var err error

	if request.RequestBody.Type == "LaunchRequest" {
		response, err = r.launchRequestHandler(request)
	} else if request.RequestBody.Type == "IntentRequest" {
		intentType := request.RequestBody.Intent.Name
		response, err = r.intentRequestHandlers[intentType](request)
	} else if request.RequestBody.Type == "SessionEndedRequest" {
		err = r.sessionEndedRequestHandler(request)
	}
	if err != nil && r.errorHandler != nil {
		return r.errorHandler(request, err)
	}
	return response
}

// AddLaunchRequestHandler sets the launch request handler funtion.
func (r *RequestHandler) AddLaunchRequestHandler(handler func(request.Request) (response.Response, error)) *RequestHandler {
	r.launchRequestHandler = handler
	return r
}

// AddIntentRequestHandler sets the intent request handler function for the given intent.
func (r *RequestHandler) AddIntentRequestHandler(intentName string, handler func(request.Request) (response.Response, error)) *RequestHandler {
	r.intentRequestHandlers[intentName] = handler
	return r
}

// AddSessionEndedRequestHandler sets the session ended request handler funtion.
func (r *RequestHandler) AddSessionEndedRequestHandler(handler func(request.Request) error) *RequestHandler {
	r.sessionEndedRequestHandler = handler
	return r
}
