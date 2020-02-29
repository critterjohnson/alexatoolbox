package request

type Request struct {
	Version     string      `json:"version,omitempty"`
	Session     Session     `json:"session,omitempty"`
	Context     Context     `json:"context,omitempty"`
	RequestBody RequestBody `json:"request,omitempty"`
}

type Session struct {
	New         bool                   `json:"new,omitempty"`
	ID          string                 `json:"sessionId,omitempty"`
	Application Application            `json:"application,omitempty"`
	Attributes  map[string]interface{} `json:"Attributes,omitempty"`
	User        User                   `json:"user,omitempty"`
}

type Application struct {
	ID string `json:"applicationId,omitempty"`
}

type User struct {
	ID          string            `json:"userId,omitempty"`
	AccessToken string            `json:"accessToken,omitempty"`
	Permissions map[string]string `json:"permissions,omitempty"`
}

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

type System struct {
	Device Device `json:"device,omitempty"`
}

type Device struct {
	ID                  string                 `json:"deviceId,omitempty"`
	SupportedInterfaces map[string]interface{} `json:"supportedInterfaces,omitempty"`
}

type Person struct {
	ID          string `json:"personId,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
}

//! read about AudioPlayer and update this
type AudioPlayer struct {
	PlayerActivity       string `json:"playerActivity,omitempty"`
	Token                string `json:"token,omitempty"`
	OffsetInMilliseconds int    `json:"offsetInMilliseconds,omitempty"`
}
