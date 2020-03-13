# go-ask
> Some tools to build Alexa skills in Go

## alexarequest
### Request object
The go-ask/alexarequest Request object is a representation of the [Alexa Request JSON](https://developer.amazon.com/en-US/docs/alexa/custom-skills/request-and-response-json-reference.html#request-format).

It is almost 1:1 (although with leading capitals, for exporting reasons), except there is no Request.Session.Attributes object. To extract session attributes, call `Request.Session.UnmarshalAttributes(obj interface{})`, which will unmarshal the session attributes to `obj` using the `encoding/json` rules.

The Request object stores the request body in the Request.RequestBody object (rather than Request.Request, as in the JSON reference) for readability purposes.

Currently, the Request object does not support interfaces. Not sure when I'll implement that, but feel free to contribute.

## alexaresponse
### Response object
The go-ask/alexaresponse Response object is a representation of the [Alexa Response JSON](https://developer.amazon.com/en-US/docs/alexa/custom-skills/request-and-response-json-reference.html#response-format).

It is 1:1, with leading capitals for names.

Response.SessionAttributes must be an `encoding/json` serializable object; same goes for Response.Response.Directives.

### Builder object
The go-ask/alexaresponse Builder object builds Alexa responses. To create a new builder, call `alexaresponse.NewBuilder()`. There are a variety of functions to add the different components of a response, make sure to [check out the code](lexaresponse/responsebuilder.go).

When you've added all you need to add, call `.Build()` to build into a go-ask/alexarespose Response object.

You don't, of course, need to use a ResponseBuilder object to build your responses, you can manually create a Response object and add to it as you please.

## handlers
### RequestHandler object
The RequestHandler object handles incoming requests. Create a new one with `NewRequestHandler()` before adding your various handlers to it. To see which handlers you can add, [check out the code](handlers/handlers.go).

When `Handle` is called with an `alexarequest.Request` object, it determines what the type of request is and passes the request along to that handler. You can also add an error handler, which, on one of your handlers erroring, will be called. That way, you can log, safe session attributes, and return a response that isn't the default error message.

The RequestHandler object can't yet handle all types of requests, and I don't know when I'll get around to it, but feel free to contribute.

## Examples
To get a feel for how to use go-ask, check out [the AWS Lambda example](examples/HelloSkill_lambda). AWS Lambda was the original intended use case, but go-ask could easily be used with another hosting platform or on your own endpoint.
