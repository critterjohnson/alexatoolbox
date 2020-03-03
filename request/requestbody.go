package request

// Body represents the body of the Request object.
type Body struct {
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

// Intent represents an intent object in an IntentRequest.
type Intent struct {
	Name               string          `json:"name,omitempty"`
	ConfirmationStatus string          `json:"confirmationStatus,omitempty"`
	Slots              map[string]Slot `json:"slots,omitempty"`
}

// Slot represents a slot in an intent.
type Slot struct {
	Name               string      `json:"name,omitempty"`
	Value              string      `json:"value,omitempty"`
	ConfirmationStatus string      `json:"confirmationStatus,omitempty"`
	Resolutions        Resolutions `json:"resolutions,omitempty"`
}

// Resolutions represents a Resolutions object.
type Resolutions struct {
	ResolutionsPerAuthority []struct {
		Authority string `json:"authority,omitempty"`
		Status    struct {
			Code string `json:"code,omitempty"`
		} `json:"status,omitempty"`
		Values []struct {
			Value ResolutionValue `json:"value,omitempty"`
		} `json:"values,omitempty"`
	} `json:"resolutionsPerAuthority,omitempty"`
}

// ResolutionValue represents a resolution value.
type ResolutionValue struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}
