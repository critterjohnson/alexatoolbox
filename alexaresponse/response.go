package alexaresponse

// Response represents the response to return from the handler.
type Response struct {
	Version           string      `json:"version,omitempty"`
	SessionAttributes interface{} `json:"sessionAttributes,omitempty"`
	Response          *Body       `json:"response,omitempty"`
}

// Body represents the body of the response to return from the handlers.
type Body struct {
	OutputSpeech     *OutputSpeech `json:"outputSpeech,omitempty"`
	Card             *Card         `json:"card,omitempty"`
	Reprompt         *Reprompt     `json:"reprompt,omitempty"`
	Directives       []interface{} `json:"directives,omitempty"`
	ShouldEndSession bool          `json:"shouldEndSession,omitempty"`
}

// Card represents the Card object.
type Card struct {
	Type    string `json:"type,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Text    string `json:"text,omitempty"`
	Image   *Image `json:"image,omitempty"`
}

// Image represents the image in a card.
type Image struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	LargeImageURL string `json:"largeImageUrl,omitempty"`
}

// OutputSpeech represents what Alexa will say.
type OutputSpeech struct {
	Type         string `json:"type,omitempty"`
	Text         string `json:"text,omitempty"`
	Ssml         string `json:"ssml,omitempty"`
	PlayBehavior string `json:"playBehavior,omitempty"`
}

// Reprompt contains an OutputSpeech object and represents what Alexa should say to reprompt the user.
type Reprompt struct {
	OutputSpeech *OutputSpeech `json:"outputSpeech,omitempty"`
}
