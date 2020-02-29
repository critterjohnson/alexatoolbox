package request

type Request struct {
	Version     string      `json:"version"`
	Session     Session     `json:"session"`
	Context     Context     `json:"context"`
	RequestBody RequestBody `json:"request"`
}

type Session struct {
	New         bool                   `json:"new"`
	ID          string                 `json:"sessionId"`
	Application Application            `json:"application"`
	Attributes  map[string]interface{} `json:"Attributes"`
	User        User                   `json:"user"`
}

type Application struct {
	ID string `json:"applicationId"`
}

type User struct {
	ID          string            `json:"userId"`
	AccessToken string            `json:"accessToken"`
	Permissions map[string]string `json:"permissions"`
}

type Context struct {
	System         System      `json:"System"`
	Device         Device      `json:"device"`
	Application    Application `json:"application"`
	User           User        `json:"user"`
	Person         Person      `json:"person"`
	APIEndpoint    string      `json:"apiEndpoint"`
	APIAccessToken string      `json:"apiAccessToken"`
	AudioPlayer    AudioPlayer `json:"AudioPlayer"`
}

type System struct {
	Device Device `json:"device"`
}

type Device struct {
	ID                  string                 `json:"deviceId"`
	SupportedInterfaces map[string]interface{} `json:"supportedInterfaces"`
}

type Person struct {
	ID          string `json:"personId"`
	AccessToken string `json:"accessToken"`
}

//! read about AudioPlayer and update this
type AudioPlayer struct {
	PlayerActivity       string `json:"playerActivity"`
	Token                string `json:"token"`
	OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
}
