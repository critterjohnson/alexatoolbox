package response

// Builder builds Alexa responses.
type Builder struct {
	response          Response
	sessionAttributes map[string]interface{}
	outputSpeech      OutputSpeech
	card              Card
	reprompt          Reprompt
	directives        []interface{}
	shouldEndSession  bool
}

func NewBuilder() *Builder {
	return &Builder{
		sessionAttributes: make(map[string]interface{}),
	}
}

// WithAttributes takes an attributes map and sets the attributes, overwriting any existing attributes.
func (b *Builder) WithAttributes(attributes map[string]interface{}) *Builder {
	b.sessionAttributes = attributes
	return b
}

// AddAttribute adds an attribute to the existing SessionAttributes map given a key and an interface.
// Adding an attribute to a key that already exist overwrites the attribute.
func (b *Builder) AddAttribute(key string, attribute interface{}) *Builder {
	b.sessionAttributes[key] = attribute
	return b
}

// AddAttributes adds all of the attributes from the current map into the existing session attributes.
// Doesn't overwrite the current session attributes unless keys are the same.
func (b *Builder) AddAttributes(attributes map[string]interface{}) *Builder {
	for k, v := range attributes {
		b.sessionAttributes[k] = v
	}
	return b
}

// WithOutputSpeech sets the OutputSpeech object of the builder to the given OutputSpeech object.
func (b *Builder) WithOutputSpeech(outputSpeech OutputSpeech) *Builder {
	b.outputSpeech = outputSpeech
	return b
}

// WithTextOutputSpeech sets the OutputSpeech type to "PlainText" and sets the text.
func (b *Builder) WithTextOutputSpeech(text string) *Builder {
	b.outputSpeech = OutputSpeech{
		Type: "PlainText",
		Text: text,
	}
	return b
}

// WithSsmlOutputSpeech sets the OutputSpeech type to "ssml" and sets the ssml encoded string.
func (b *Builder) WithSsmlOutputSpeech(ssml string) *Builder {
	b.outputSpeech = OutputSpeech{
		Type: "ssml",
		Ssml: ssml, //TODO: Create an ssml builder
	}
	return b
}

// OutputSpeechPlayBehavior sets the play behavior of the output speech.
func (b *Builder) OutputSpeechPlayBehavior(behavior string) *Builder {
	b.outputSpeech.PlayBehavior = behavior
	return b
}

// WithCard sets the Card object of the builder to the given Card object.
func (b *Builder) WithCard(card Card) *Builder {
	b.card = card
	return b
}

// WithSimpleCard creates a simple card.
func (b *Builder) WithSimpleCard(title string, content string) *Builder {
	b.card = Card{
		Type:    "Simple",
		Title:   title,
		Content: content,
	}
	return b
}

// WithStandardCard creates a standard card.
func (b *Builder) WithStandardCard(title string, text string, image Image) *Builder {
	b.card = Card{
		Title: title,
		Text:  text,
		Image: image,
	}
	return b
}

// WithReprompt sets the reprompt object of the builder to the given reprompt object.
func (b *Builder) WithReprompt(reprompt Reprompt) *Builder {
	b.reprompt = reprompt
	return b
}

// WithTextReprompt creates the reprompt object with a "PlainText" type OutputSpeech.
func (b *Builder) WithTextReprompt(text string) *Builder {
	b.reprompt = Reprompt{
		OutputSpeech: OutputSpeech{
			Type: "PlainText",
			Text: text,
		},
	}
	return b
}

// WithSsmlReprompt creates the reprompt object with a "ssml" type OutputSpeech.
func (b *Builder) WithSsmlReprompt(ssml string) *Builder {
	b.reprompt = Reprompt{
		OutputSpeech: OutputSpeech{
			Type: "ssml",
			Ssml: ssml,
		},
	}
	return b
}

// RepromptPlayBehavior sets the play behavior of the reprompt.
func (b *Builder) RepromptPlayBehavior(behavior string) *Builder {
	b.reprompt.OutputSpeech.PlayBehavior = behavior
	return b
}

// WithDirectives sets the directives for the response.
func (b *Builder) WithDirectives(directives []interface{}) *Builder {
	b.directives = directives
	return b
}

// ShouldEndSession determines whether the session should end after this request.
func (b *Builder) ShouldEndSession(boolean bool) *Builder {
	b.shouldEndSession = boolean
	return b
}

// Build builds and returns the Response object.
func (b *Builder) Build() Response {
	return Response{
		Version:           "1.0",
		SessionAttributes: b.sessionAttributes,
		Response: ResponseBody{
			OutputSpeech:     b.outputSpeech,
			Card:             b.card,
			Reprompt:         b.reprompt,
			Directives:       b.directives,
			ShouldEndSession: b.shouldEndSession,
		},
	}
}
