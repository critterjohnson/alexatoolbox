package request

type Intent struct {
	Name               string          `json:"name"`
	ConfirmationStatus string          `json:"confirmationStatus"`
	Slots              map[string]Slot `json:"slots"`
}

type Slot struct {
	Name               string      `json:"name"`
	Value              string      `json:"value"`
	ConfirmationStatus string      `json:"confirmationStatus"`
	Resolutions        Resolutions `json:"resolutions"`
}

type Resolutions struct {
	ResolutionsPerAuthority []struct {
		Authority string `json:"authority"`
		Status    struct {
			Code string `json:"code"`
		} `json:"status"`
		Values []struct {
			Value struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"value"`
		} `json:"values"`
	} `json:"resolutionsPerAuthority"`
}
