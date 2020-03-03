package handlers

import (
	request "github.com/critterjohnson/go-ask/alexarequest"
	response "github.com/critterjohnson/go-ask/alexaresponse"
)

// RequestHandler handles incoming requests and sends them to user-defined handlers.
type RequestHandler struct {
	launchRequestHandler       func(request.Request) (response.Response, error)
	intentRequestHandlers      map[string]func(request.Request) (response.Response, error)
	sessionEndedRequestHandler func(request.Request) error
	errorHandler               func(request.Request, error) (response.Response, error)
}

// NewRequestHandler returns a new request handler.
func NewRequestHandler() *RequestHandler {
	return &RequestHandler{
		intentRequestHandlers: make(map[string]func(request.Request) (response.Response, error)),
	}
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

// AddErrorHandler adds a function to be called if an error occurs. Useful for logging.
func (r *RequestHandler) AddErrorHandler(handler func(request.Request, error) (response.Response, error)) *RequestHandler {
	r.errorHandler = handler
	return r
}

// Handle handles an incoming request by calling the user-defined handlers.
func (r *RequestHandler) Handle(request request.Request) (response.Response, error) {
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
	return response, err
}
