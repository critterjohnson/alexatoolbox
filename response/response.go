package response

type Response struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes"`
	Response          ResponseBody           `json:"response"`
}

type ResponseBody struct {
	OutputSpeech     OutputSpeech  `json:"outputSpeech"`
	Card             Card          `json:"card"`
	Reprompt         Reprompt      `json:"reprompt"`
	Directives       []interface{} `json:"directives"`
	ShouldEndSession bool          `json:"shouldEndSession"`
}

type Card struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Text    string `json:"text"`
	Image   Image  `json:"image"`
}

type Image struct {
	SmallImageURL string `json:"smallImageUrl"`
	LargeImageURL string `json:"largeImageUrl"`
}

type OutputSpeech struct {
	Type         string `json:"type"`
	Text         string `json:"text"`
	Ssml         string `json:"ssml"`
	PlayBehavior string `json:"playBehavior"`
}

type Reprompt struct {
	OutputSpeech OutputSpeech `json:"outputSpeech"`
}
