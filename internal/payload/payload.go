package payload

import "encoding/json"

type ButtonCommand string

const (
	StartCommand ButtonCommand = "start"
	SexCommand   ButtonCommand = "sex"
	AboutCommand ButtonCommand = "about"
	CityCommand  ButtonCommand = "city"
	AgeCommand   ButtonCommand = "age"
	SaveCommand  ButtonCommand = "save"
)

type Payload struct {
	Command ButtonCommand `json:"command"`
	Options Options       `json:"options,omitempty"`
}

type Options struct {
	InterestedIn string `json:"interested_in,omitempty"`
	PhotoURL     string `json:"photo_url,omitempty"`
	Name         string `json:"name,omitempty"`
	Age          int    `json:"age,omitempty"`
	City         string `json:"city,omitempty"`
}

func New(command ButtonCommand) *Payload {
	return &Payload{
		Command: command,
	}
}

func Unmarshal(rawPayload string) (payload Payload, err error) {
	err = json.Unmarshal([]byte(rawPayload), &payload)
	return
}
