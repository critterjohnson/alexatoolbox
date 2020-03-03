package handlers

import (
	"errors"
	"testing"

	"github.com/critterjohnson/go-ask/request"
	"github.com/critterjohnson/go-ask/response"
	"github.com/critterjohnson/go-ask/testgen"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewRequestHandler(t *testing.T) {
	Convey("When NewRequestHandler is called", t, func() {
		handle := NewRequestHandler()

		Convey("it should return the RequestHandler object", func() {
			So(handle, ShouldNotBeNil)
		})
	})
}

func TestAddLaunchRequestHandler(t *testing.T) {
	handle := NewRequestHandler()

	Convey("When AddLaunchRequestHandler is called", t, func() {
		launchHandler := func(request request.Request) (response.Response, error) {
			return response.Response{}, nil
		}
		handle.AddLaunchRequestHandler(launchHandler)

		Convey("it should pass the handler function", func() {
			So(handle.launchRequestHandler, ShouldEqual, launchHandler)
		})
	})
}

func TestAddntentRequestHandler(t *testing.T) {
	handle := NewRequestHandler()

	Convey("When AddIntentRequestHandler is called", t, func() {
		intentRequestHandler := func(request request.Request) (response.Response, error) {
			return response.Response{}, nil
		}
		handle.AddIntentRequestHandler("intentName", intentRequestHandler)

		Convey("it should add the handler function", func() {
			So(handle.intentRequestHandlers["intentName"], ShouldEqual, intentRequestHandler)
		})
	})
}

func TestSessionEndedRequestHandler(t *testing.T) {
	handle := NewRequestHandler()

	Convey("When AddSessionEndedRequestHandler is called", t, func() {
		sessionEndedHandler := func(request request.Request) error {
			return nil
		}
		handle.AddSessionEndedRequestHandler(sessionEndedHandler)

		Convey("it should pass the handler function", func() {
			So(handle.sessionEndedRequestHandler, ShouldEqual, sessionEndedHandler)
		})
	})
}

func TestAddErrorHandler(t *testing.T) {
	handle := NewRequestHandler()

	Convey("When AddErrorHandler is called", t, func() {
		errorHandler := func(request request.Request, err error) (response.Response, error) {
			return response.Response{}, nil
		}
		handle.AddErrorHandler(errorHandler)

		Convey("it should pass the handler function", func() {
			So(handle.errorHandler, ShouldEqual, errorHandler)
		})
	})
}

func TestHandle(t *testing.T) {
	handle := NewRequestHandler()

	Convey("When Handle is called", t, func() {
		req := request.RandomRequest(t)

		Convey("on a LaunchRequest", func() {
			req.RequestBody.Type = "LaunchRequest"
			e := errors.New(testgen.RandomString())
			res := response.RandomResponse(t)

			launchReqHandler := func(launchRequest request.Request) (response.Response, error) {
				Convey("it should pass the request", func() {
					So(launchRequest, ShouldResemble, req)
				})

				return res, e
			}
			handle.AddLaunchRequestHandler(launchReqHandler)
			response, err := handle.Handle(req)

			Convey("it should return the response", func() {
				So(response, ShouldResemble, res)
			})

			Convey("it should return the error", func() {
				So(err, ShouldEqual, e)
			})
		})

		Convey("on an IntentRequest", func() {
			req.RequestBody.Type = "IntentRequest"
			intentName := testgen.RandomString()
			req.RequestBody.Intent.Name = intentName

			e := errors.New(testgen.RandomString())
			res := response.RandomResponse(t)

			launchReqHandler := func(intentRequest request.Request) (response.Response, error) {
				Convey("it should pass the request", func() {
					So(intentRequest, ShouldResemble, req)
				})

				return res, e
			}
			handle.AddIntentRequestHandler(intentName, launchReqHandler)
			response, err := handle.Handle(req)

			Convey("it should return the response", func() {
				So(response, ShouldResemble, res)
			})

			Convey("it should return the error", func() {
				So(err, ShouldEqual, e)
			})
		})

		Convey("on a SessionEnded", func() {
			req.RequestBody.Type = "SessionEndedRequest"
			e := errors.New(testgen.RandomString())

			endReqHandler := func(endRequest request.Request) error {
				Convey("it should pass the request", func() {
					So(endRequest, ShouldResemble, req)
				})

				return e
			}
			handle.AddSessionEndedRequestHandler(endReqHandler)
			resp, err := handle.Handle(req)

			Convey("it should return an empty response", func() {
				So(resp, ShouldResemble, response.Response{})
			})

			Convey("it should return the error", func() {
				So(err, ShouldEqual, e)
			})
		})
	})
}
