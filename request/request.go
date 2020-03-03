package request

// Request represents the request sent by Alexa.
type Request struct {
	Version     string  `json:"version,omitempty"`
	Session     Session `json:"session,omitempty"`
	Context     Context `json:"context,omitempty"`
	RequestBody Body    `json:"request,omitempty"`
}

// Session represents the session object.
type Session struct {
	New         bool                   `json:"new,omitempty"`
	ID          string                 `json:"sessionId,omitempty"`
	Application Application            `json:"application,omitempty"`
	Attributes  map[string]interface{} `json:"Attributes,omitempty"`
	User        User                   `json:"user,omitempty"`
}

// Application represents the Application object.
type Application struct {
	ID string `json:"applicationId,omitempty"`
}

// User represent's the skill's user.
type User struct {
	ID          string            `json:"userId,omitempty"`
	AccessToken string            `json:"accessToken,omitempty"`
	Permissions map[string]string `json:"permissions,omitempty"`
}

// Context represents the context of the request.
type Context struct {
	System         System      `json:"System,omitempty"`
	Device         Device      `json:"device,omitempty"`
	Application    Application `json:"application,omitempty"`
	User           User        `json:"user,omitempty"`
	Person         Person      `json:"person,omitempty"`
	APIEndpoint    string      `json:"apiEndpoint,omitempty"`
	APIAccessToken string      `json:"apiAccessToken,omitempty"`
	AudioPlayer    AudioPlayer `json:"AudioPlayer,omitempty"`
}

// System represents the system making the request.
type System struct {
	Device Device `json:"device,omitempty"`
}

// Device represents the device making the request.
type Device struct {
	ID                  string                 `json:"deviceId,omitempty"`
	SupportedInterfaces map[string]interface{} `json:"supportedInterfaces,omitempty"`
}

// Person represents the person making the request.
type Person struct {
	ID          string `json:"personId,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
}

// AudioPlayer represents the AudioPlayer object. Not properly implemented.
type AudioPlayer struct {
	PlayerActivity       string `json:"playerActivity,omitempty"`
	Token                string `json:"token,omitempty"`
	OffsetInMilliseconds int    `json:"offsetInMilliseconds,omitempty"`
} //! read about AudioPlayer and update this
