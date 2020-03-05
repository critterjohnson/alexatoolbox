# Example - Using go-ask with AWS Lambda
Build a skill with AWS Lambda serverless compute and Go.

## Getting Started
### AWS Account
[Make an AWS Account.](https://portal.aws.amazon.com/billing/signup?redirect_url=https%3A%2F%2Faws.amazon.com%2Fregistration-confirmation#/start) Note that this does require payment information, but nothing we do here will be taking outside of the AWS Fre Tier.

### Amazon Account
[Make an Amazon account](https://www.amazon.com/ap/register?_encoding=UTF8&openid.assoc_handle=usflex&openid.claimed_id=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.identity=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.mode=checkid_setup&openid.ns=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0&openid.ns.pape=http%3A%2F%2Fspecs.openid.net%2Fextensions%2Fpape%2F1.0&openid.pape.max_auth_age=0&openid.return_to=https%3A%2F%2Fwww.amazon.com%2Fgp%2Fyourstore%2Fhome%3Fie%3DUTF8%26ref_%3Dnav_newcust), if you don't already have one. This is what we'll be using for building our Alexa Skill.

## AWS Lambda
AWS Lambda is Amazon Web Services' "serverless cloud compute" service. It abstracts the developer from the server, and charges you only for your compute time. You set up individual functions that respond to an event; it can respond to other AWS Services (such as the Alexa skills kit!) or be called directly. Check out https://aws.amazon.com/lambda/ for more info.

### Our first Lambda function
On the AWS Console, click on the "Services" dropdown at the top left and on "Lambda" under "Compute". You should be taken to a screen that looks like this:
![dashboard.png][dashboard]
Once you're here, click "Create function." You'll be taken to the create function screen, where you should click "Start from scratch", name your function "GoHello", set your runtime to "Go 1.x", and finally click Create function.
![create_function.png][create_function]
Note the "Permissions" tab - we won't be messing around with it now, but if you wanted your code to be allowed to access other AWS services this is where you'd set that up.

Scroll down to "Function code" and change the Handler to `main`. We won't be writing our code just yet, but this will make sense later.

## Alexa Skills
Alexa skills are just apps for Alexa, and are essentially just a bunch of interactions between Alexa and some server. Alexa sends a request to an endpoint we set, and our endpoint has to respond to it in a certain way. We can control what she says, the sounds she plays, display an image (for devices that support it), and a whole lot more. We'll be going through the process to make a very basic skill, and hopefully picking up the essentials along the way. If anything confuses you, be sure to check out the docs or the [glossary](https://developer.amazon.com/en-US/alexa/alexa-skills-kit/resources/key-terms).

### The flow
The basic flow of an Alexa skill is very simple. With a couple exceptions, it looks like this:
1. The user says something, Alexa parses the what the user says into some JSON.
2. Alexa sends the JSON request to an endpoint, along with some other useful information, such as what types of responses the device supports (display images, play videos, etc).
3. The service at the endpoint receives the JSON data. It's up to the serivice to use the JSON data to figure out what type of request it is (more on that later), and process the data accordingly.
4. The service at the endpoint creates a response object based off of the request data, and sends it back to Alexa.
5. Alexa does whatever the response tells her to do.

Some notable exceptions:
* When the user ends the session, a request is sent to the endpoint, but the endpoint cannot send a response. This is useful for logging or saving session data.
* When a user does something to control the playback of what they're listening to - pausing it, playing it, etc - a request is sent to the endpoint.

### What do we have control over?
The skill developer has control over almost every step of the process. The designer defines:
1. How Alexa should expect requests to be made (and therefore how to parse the speech)
2. The endpoint to make requests to
3. The code at the endpoint, how responses are generated

### What does the request JSON object look like?
From the [Request and Response JSON Reference]()
```json
{
  "version": "1.0",
  "session": {
    "new": true,
    "sessionId": "amzn1.echo-api.session.[unique-value-here]",
    "application": {
      "applicationId": "amzn1.ask.skill.[unique-value-here]"
    },
    "attributes": {
      "key": "string value"
    },
    "user": {
      "userId": "amzn1.ask.account.[unique-value-here]",
      "accessToken": "Atza|AAAAAAAA...",
      "permissions": {
        "consentToken": "ZZZZZZZ..."
      }
    }
  },
  "context": {
    "System": {
      "device": {
        "deviceId": "string",
        "supportedInterfaces": {
          "AudioPlayer": {}
        }
      },
      "application": {
        "applicationId": "amzn1.ask.skill.[unique-value-here]"
      },
      "user": {
        "userId": "amzn1.ask.account.[unique-value-here]",
        "accessToken": "Atza|AAAAAAAA...",
        "permissions": {
          "consentToken": "ZZZZZZZ..."
        }
      },
      "person": {
        "personId": "amzn1.ask.account.[unique-value-here]",
        "accessToken": "Atza|BBBBBBB..."
      },
      "apiEndpoint": "https://api.amazonalexa.com",
      "apiAccessToken": "AxThk..."
    },
    "AudioPlayer": {
      "playerActivity": "PLAYING",
      "token": "audioplayer-token",
      "offsetInMilliseconds": 0
    }
  },
  "request": {}
}
```
Check out the reference to see what all of this means, but there are a few key things I should point out:
#### Session Attributes
```json
"session" {
    "attributes": {}
}
```
Session Attributes are the only way to have persistence within a session. Because there is no master process always running, and our code can only receive and respond to requests, this is where information must be stored so that our service knows what the user has already said. We won't be using it, but it's important to know about.

#### Request Object
```json
"request": {}
```
The Request object, or the Request Body, is where the actual request is stored. There are several request types, each with their own properties:
* LaunchRequest
    * Required - sent when the user starts the skill with no other information. For our skill, that would look like: "Alexa, start Go Hello."
* CanFulfillIntentRequest
    * Optional - sent when Alexa isn't sure which skill the user is trying to hit.
* IntentRequest
    * Required - this is where the main functionality of the skill comes in. This is sent when the user is trying to do something with our skill. See Intents below.
* SessionEndedRequest
    * Required - sent when the user ends the interraction. The skill cannot respond to this.

They each have their own format, so your code needs to be prepared to handle each one. There are even more types for different interfaces. Check out the [Request Types Reference](https://developer.amazon.com/en-US/docs/alexa/custom-skills/request-types-reference.html) for further learning.

## Making a skill using go-ask
### Setup
Sign into the [Alexa Skills Developer Console](https://developer.amazon.com/alexa/console/ask).
![dev_console.png][dev_console]
Click "Create Skill", name it "GoHello", and leave the defaults ("Custom" and "Provision your own"). Click "Create skill", and on the next screen click "Start from scratch" and "Choose" at the top right. You'll be taken to the Skill Dashboard (mine will has a few more skills than yours will):
![skill_dashboard.png][skill_dashboard]
Under "Invocation" on the left, change the Skill Invocation Name to "go hello". This is how we'll be reffering to our app from Alexa when we want to start it. Click "Save Model" at the top.


### Intents
On the left, click on Intents. You'll be greeted with a list of 5 intents:
* AMAZON.FallbackIntent
* AMAZON.CancelIntent
* AMAZON.HelpIntent
* AMAZON.StopIntent
* AMAZON.NavigateHomeIntent
Note that all but FallBackIntent are required. We won't be handling all of these Intents with our code, which means you couldn't take this skill public until you did.

You can think of Intents like functions. This is where we define what our skill can do. The different functionalities of our skill get separated into different events - for example, were you to make a skill to manage a to-do list, you would have an intent to create an item, remove an item, check off an item, edit an item, etc.

Our skill will have one intent that will say hello to a user. When they say "Alexa, tell Go Hello my name is {name}", it will respond "Hello, {name}" (with whatever name the user gave it).

Let's create our intent. Click "Add intent", and name your intent "SayHello".
![create_intent.png][create_intent]
The next screen you're on will allow you to create "sample utterances". When the user talks, Alexa will try to fit their speech to a sample utterance, filling in the necessary slots. Slots are like variables that are "filled" by the user's voice. Slots have types, which are a list of words Alexa should expect to hear there. When the user says something that fits a sample utterance from an intent, Alexa knows which intent it is, fills the slots, and ships that request off to our endpoint. This will make more sense once we've gotten started on our sample utterances.

Begin with the sample utterance `my name is {name}`. The braces are our slot, that we've decided to call "name". Once the utterance is created, down below, under Intent Slots, you will see a slot named "name". On the dropdown next to it, click and type "AMAZON.US_FIRST_NAME". This is a built in list of names. You can, of course, define your own slot types, but we won't need to do that here since the list is premade.
![intent_screen.png][intent_screen]
Click "Save Model" at the top.

Good practice is to include the different ways the user might be invoking this intent - for instance, you could also add `hello, my name is {name}` to cover that possible way of invoking our intent.

When the user says "Alexa, tell GoHello my name is Christopher", Alexa will parse out the speech into a structure that looks like this (modified from the [Request Types Reference](https://developer.amazon.com/en-US/docs/alexa/custom-skills/request-types-reference.html)):
```json
"request": {
    "type": "IntentRequest",
    "requestId": "string",
    "timestamp": "string",
    "dialogState": "string",
    "locale": "string",
    "intent": {
        "name": "SayHello",
        "confirmationStatus": "string",
        "slots": {
        "name": {
            "name": "name",
            "value": "Christopher",
            "confirmationStatus": "string",
            "resolutions": {
            "resolutionsPerAuthority": [
                {
                "authority": "string",
                "status": {
                    "code": "string"
                },
                "values": [
                    {
                    "value": {
                        "name": "string",
                        "id": "string"
                    }
                    }
                ]
                }
            ]
            }
        }
        }
    }
}
```
All our code will have to handle is what kind of request it is (in this case, an IntentRequest), the intent (SayHello), and the value of the slots (we have one, called "name", with the value "Christopher", in this case). First, we have to tell Alexa where to find our code.


### Setting our endpoint
On the far left, click "Endpoint." We're going to do some jumping around here between this screen and the AWS Lambda screen, so bear with me. Click "AWS Lambda ARN". Copy the skill ID to your clipboard. Back on the Lambda screen, click "Add Trigger". Triggers are the events that can call your AWS Lambda function. In the dropdown, click Alexa Skills Kit and paste your Skill ID.
![triggers.png][triggers]
Now that we're back on the main configuration menu, copy your ARN at the top right of the screen. This is just the ID of the lambda function. Head back over to the Alexa Developer Console screen, and paste your ARN in the "Default Region" box.
![set_arn.png][set_arn]
Click "Save Endpoints" at the top. We're officially done configuring our skill, and can *finally* get to the code.

### The Code
#### The Response object
From the [Alexa Skills Request and Response JSON Reference](https://developer.amazon.com/en-US/docs/alexa/custom-skills/request-and-response-json-reference.html#response-format)
```json
{
  "version": "string",
  "sessionAttributes": {
    "key": "value"
  },
  "response": {
    "outputSpeech": {
      "type": "PlainText",
      "text": "Plain text string to speak",
      "playBehavior": "REPLACE_ENQUEUED"      
    },
    "card": {
      "type": "Standard",
      "title": "Title of the card",
      "text": "Text content for a standard card",
      "image": {
        "smallImageUrl": "https://url-to-small-card-image...",
        "largeImageUrl": "https://url-to-large-card-image..."
      }
    },
    "reprompt": {
      "outputSpeech": {
        "type": "PlainText",
        "text": "Plain text string to speak",
        "playBehavior": "REPLACE_ENQUEUED"             
      }
    },
    "directives": [
      {
        "type": "InterfaceName.Directive"
        (...properties depend on the directive type)
      }
    ],
    "shouldEndSession": true
  }
}
```
Our code will have to create a response object that looks like this. go-ask provides a builder object to make it for you, but you do need some knowledge of what this looks like. A few notable things:
```json
"outputSpeech": {
      "type": "PlainText",
      "text": "Plain text string to speak",
      "playBehavior": "REPLACE_ENQUEUED"      
}
```
This object tells Alexa what to say. The type can be either `PlainText` or `ssml`, PlainText is just a string of words to say and ssml is marked up with ssml to have more dynamic speech.
```json
"shouldEndSession": true
```
Determines if the session should be ended on this response. A "session" is started when the user invokes our skill, which they can do several ways. The ones we have to be concerned about are:
* Saying "Alexa, start Go Hello", which will initiate a LaunchRequest. After a LaunchRequest, we want to keep our session open, so the user doesn't have to say "Alexa" again to get her to start listening. They should instead jump right into saying "my name is christopher", which will send an intent request just like the user had started the skill by saying
* "Alexa, tell Go Hello my name is Christopher" - our skill should respond and close the session.

#### Our code
Our code will be hosted on AWS Lambda. Make a new folder to store our files, and in it make a new file called `main.go`. I would suggest initializing go modules from here, with
```bash
go mod init
```
In `main.go`, add the following code:
```go
package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/critterjohnson/go-ask/alexarequest"
	"github.com/critterjohnson/go-ask/alexaresponse"
	"github.com/critterjohnson/go-ask/handlers"
)

// Launch is the handler for LaunchRequests.
func Launch(request alexarequest.Reqeust) (alexaresponse.Response, error) {
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
```

#### Walking through the code
Our import statement,
```go
import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/critterjohnson/go-ask/alexarequest"
	"github.com/critterjohnson/go-ask/alexaresponse"
	"github.com/critterjohnson/go-ask/handlers"
)
```
imports a few things for us:
* `context` - the Golang built-in context library for managing the context of our application.
* `github.com/aws/aws-lambda-go/lambda` - the AWS library for managing lambda functions. The [Lambda documentation](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html) will explain this much better than I can, you should check it out.
* `github.com/critterjohnson/go-ask/alexarequest` - the go-ask request package that represents the Alexa skill request JSON object.
* `github.com/critterjohnson/go-ask/alexaresponse` - the go-ask response package that represents an Alexa JSON response object and provides a tool for building responses.
* `github.com/critterjohnson/go-ask/handlers` - the go-ask package that allows you to create handlers for individual request types and pass them to one master Handler object.

The LaunchRequest handler function:
```go
func Launch(request alexarequest.Reqeust) (alexaresponse.Response, error) {
	return alexaresponse.NewBuilder().
		WithTextOutputSpeech("Welcome to GoHello! Please say hello and introduce yourself with your name.").
		ShouldEndSession(false).
		Build(), nil
}
```
creates a new response builder, adds some output speech, **sets shouldEndSession to false**, and builds the response.

The SayHello function
```go
func SayHello(request alexarequest.Request) (alexaresponse.Response, error) {
	return alexaresponse.NewBuilder().
		WithTextOutputSpeech("Hello, " + request.RequestBody.Intent.Slots["name"].Value).
		Build(), nil
}
```
is the function we'll use to handle incoming IntentRequests with the SayHello intent. We set the outputSpeech to "Hello, " and concatenate the value of the "name" slot.

Finally, End
```go
func End(request alexarequest.Request) error {
	return nil
}
```
will handle our SessionEndedRequests.

Our LambdaHandler function
```go
func LambdaHandler(ctx context.Context, request alexarequest.Request) (alexaresponse.Response, error) {
	requestHandler := handlers.NewRequestHandler()
	requestHandler.AddLaunchRequestHandler(Launch)
	requestHandler.AddIntentRequestHandler("SayHello", SayHello)
	requestHandler.AddSessionEndedRequestHandler(End)
	response, err := requestHandler.Handle(request)
	return response, err
}
```
Creates a new request handler, and adds our handler functions to it. You can see that on the line 
```go
requestHandler.AddIntentRequestHandler("SayHello", SayHello)
```
we're telling the Handler that our SayHello function will handle "SayHello" intents.
requestHandler.Handle handles the incoming alexarequest.Request with whatever function we've set to handle that function type.

Lastly,
```go
func main() {
	lambda.Start(LambdaHandler)
}
```
is our main function that sets the Lambda Handler function.

## Uploading our code to Lambda
AWS Lambda runs on Linux servers. Go is a compiled language, so you'll have to compile for Linux. Luckily, Go has its own built-in way of doing this. Run the command
```bash
env GOOS=linux GOARCH=amd64 go build -v -o bin/main
```
to build for Linux. Note `-o bin/main`, which will output our executable in a new directory `bin` and called `main` (without a file extension).

I've had problems with this in the past, sometimes you'll get an unusual "permission denied" or "PathError" error when you build with `GOOS=linux`. If that happens, spin up an Ubuntu VM and build it there. Hopefully this gets patched soon.

We're going to create a deployment package, which is just a zip file with our binary in it. Run:
```bash
zip -r main.zip bin/main
```
This will create a zip file with our binary in it and call it `main.zip`.

On the Lambda page, click "Upload" under "Upload a .zip file".
![upload_package.png][upload_package]

Click "Save" at the top right.

## Build and test our skill
Finally, the time has come. Back on the Skills Developer Console screen, click "Build Model" at the top and wait for it to complete. It might take a while.
![build_model.png][build_model]

Click "Test" at the top of your screen. On the "Skill testing is enabled in:" dropdown, click development.
![enable_testing.png][enable_testing]

In the "Type or click and hold the mic" box, type "tell go hello my name is John." It should respond "hello, John" - congratulations! You just made your first Alexa skill in Go!


[dashboard]: images/dashboard.png
[create_function]: images/create_function.png
[dev_console]: images/dev_console.png
[create_skill]: images/create_skill.png
[skill_dashboard]: images/skill_dashboard.png
[create_intent]: images/create_intent.png
[intent_screen]: images/intent_screen.png
[triggers]: images/triggers.png
[set_arn]: images/set_arn.png
[upload_package]: images/upload_package.png
[build_model]: images/build_model.png
[enable_testing]: images/enable_testing.png
