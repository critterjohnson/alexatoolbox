package request

type RequestBody struct {
	Type      string `json:"type"`
	RequestID string `json:"requestId"`
	Timestamp string `json:"timestamp"`
	Locale    string `json:"locale"`
	// IntentRequest
	DialogeState string `json:"dialogueState"`
	Intent       Intent `json:"intent"`
	// SessionEndedRequest
	Reason string `json:"reason"`
	Error  struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"error"`
}
