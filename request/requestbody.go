package request

type RequestBody struct {
	Type      string `json:"type,omitempty"`
	RequestID string `json:"requestId,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Locale    string `json:"locale,omitempty"`
	// IntentRequest
	DialogeState string `json:"dialogueState,omitempty"`
	Intent       Intent `json:"intent,omitempty"`
	// SessionEndedRequest
	Reason string `json:"reason,omitempty"`
	Error  struct {
		Type    string `json:"type,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"error,omitempty"`
}
