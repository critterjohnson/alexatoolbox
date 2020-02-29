package request

type Intent struct {
	Name               string          `json:"name,omitempty"`
	ConfirmationStatus string          `json:"confirmationStatus,omitempty"`
	Slots              map[string]Slot `json:"slots,omitempty"`
}

type Slot struct {
	Name               string      `json:"name,omitempty"`
	Value              string      `json:"value,omitempty"`
	ConfirmationStatus string      `json:"confirmationStatus,omitempty"`
	Resolutions        Resolutions `json:"resolutions,omitempty"`
}

type Resolutions struct {
	ResolutionsPerAuthority []struct {
		Authority string `json:"authority,omitempty"`
		Status    struct {
			Code string `json:"code,omitempty"`
		} `json:"status,omitempty"`
		Values []struct {
			Value struct {
				Name string `json:"name,omitempty"`
				ID   string `json:"id,omitempty"`
			} `json:"value,omitempty"`
		} `json:"values,omitempty"`
	} `json:"resolutionsPerAuthority,omitempty"`
}
