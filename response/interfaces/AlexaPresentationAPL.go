package interfaces

type Document struct {
	Commands      map[string]interface{}            `json:"command"`
	Description   string                            `json:"description"`
	Graphics      map[string]AVG                    `json:"graphics"`
	HandleKeyDown []KeyHandler                      `json:"handleKeyDown"`
	HandleKeyUP   []KeyHandler                      `json:"handleKeyUp"`
	Imports       []Import                          `json:"imports"`
	Layouts       map[string]Layout                 `json:"layouts"`
	MainTemplate  Layout                            `json:"mainTemplate"`
	OnMount       []interface{}                     `json:"onMount"`
	Resources     []Resource                        `json:"resources"`
	Settings      map[string]map[string]interface{} `json:"settings"`
	Styles        map[string]string                 `json:"styles"`
	Theme         string                            `json:"theme"`
	Type          string                            `json:"type"`
	Version       string                            `json:"version"`
}

type AVG struct {
	Description     string        `json:"description"`
	Height          string        `json:"height"`
	Items           []interface{} `json:"items"`
	Parameters      []Parameter   `json:"parameters"`
	ScaleTypeHeight string        `json:"scaleTypeHeight"`
	ScaleTypeWidth  string        `json:"scaleTypeWidth"`
	Type            string        `json:"type"`
	Version         string        `json:"version"`
	ViewportHeight  float32       `json:"viewportHeight"`
	ViewportWidth   float32       `json:"viewportWidth"`
	Width           string        `json:"width"`
}

type PathItem struct {
	Type          string  `json:"path"`
	FillOpacity   float32 `json:"fillOpacity"`
	Fill          string  `json:"fillColor"`
	PathData      string  `json:"pathData"`
	StrokeOpacity float32 `json:"strokeOpacity"`
	Stroke        string  `json:"stroke"`
	StrokeWidth   float32 `json:"strokeWidth"`
}

type GroupItem struct {
	Type       string        `json:"type"`
	Opacity    float32       `json:"opacity"`
	Rotation   float32       `json:"rotation"`
	PivotX     float32       `json:"pivotX"`
	PivotY     float32       `json:"pivotY"`
	ScaleX     float32       `json:"scaleX"`
	ScaleY     float32       `json:"scaleY"`
	TranslateX float32       `json:"translateX"`
	TranslateY float32       `json:"translateY"`
	Items      []interface{} `json:"items"`
}

type KeyHandler struct {
	Commands  []interface{} `json:"commands"`
	Propagate bool          `json:"propagate"`
	When      bool          `json:"propagate"`
}

type Import struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Source  string `json:"source"`
}

type Layout struct {
	Description string      `json:"description"`
	Item        Component   `json:"item"`
	Items       []Component `json:"items"`
	Parameters  []Parameter `json:"parameters"`
}

type Component struct {
	AccessibilityLabel     string        `json:"accessibilityLabel"`
	Bind                   []Binding     `json:"bind"`
	Description            string        `json:"description"`
	Checked                bool          `json:"checked"`
	Disabled               bool          `json:"disabled"`
	Display                string        `json:"display"`
	Entity                 interface{}   `json:"entity"`
	Entities               []interface{} `json:"entities"`
	Height                 string        `json:"height"`
	ID                     string        `json:"id"`
	InheritParentState     bool          `json:"inheritParentState"`
	MaxHeight              string        `json:"maxHeight"`
	MaxWidth               string        `json:"maxWidth"`
	MinHeight              string        `json:"minHeight"`
	MinWidth               string        `json:"minWidth"`
	OnMount                []interface{} `json:"onMount"`
	OnCursorEnter          []interface{} `json:"onCursorEnter"`
	OnCursorExit           []interface{} `json:"onCursorExit"`
	Opacity                float32       `json:"opacity"`
	PaddingBottom          string        `json:"paddingBottom"`
	PaddingLeft            string        `json:"paddingLeft"`
	PaddingRight           string        `json:"paddingRight"`
	PaddingTop             string        `json:"paddingTop`
	ShadowColor            string        `json:"shadowColor"`
	ShadorHorizontalOffset string        `json:"shadowHorizontalOffset"`
	ShadowRadius           string        `json:"shadowRadius"`
	ShadowVerticalOffset   string        `json:"shadowVerticalOffset"`
	Speech                 interface{}   `json:"speech"`
	Style                  string        `json:"style"`
	Transform              []Transform   `json:"transform"`
	Type                   string        `json:"type"`
	When                   bool          `json:"when"`
	Width                  string        `json:"width"`
}

type Binding struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Transform struct {
	Rotate     float32 `json:"rotate"`
	Scale      float32 `json:"scale"`
	ScaleX     float32 `json:"scaleX"`
	ScaleY     float32 `json:"scaleY"`
	SkewX      float32 `json:"skewX"`
	SkewY      float32 `json:"skewY"`
	TranslateX string  `json:"translateX"`
	TranslateY string  `json:"translateY`
}

type Parameter struct {
	Default     string `json:"default"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type Resource struct {
	Boolean     map[string]bool   `json:"boolean"`
	Colors      map[string]string `json:"colors"`
	Description string            `json:"description"`
	Dimensions  map[string]string `json:"dimensions"`
	Strings     map[string]string `json:"strings"`
	When        bool              `json:"when"`
}
