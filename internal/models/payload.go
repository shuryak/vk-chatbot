package models

import "encoding/json"

type ButtonCommand string

const (
	StartCommand ButtonCommand = "start"
	SexCommand   ButtonCommand = "sex"
	AboutCommand ButtonCommand = "about"
	NameCommand  ButtonCommand = "name"
	CityCommand  ButtonCommand = "city"
	AgeCommand   ButtonCommand = "age"
	SaveCommand  ButtonCommand = "save"
)

type Payload struct {
	Command ButtonCommand  `json:"command"`
	Options PayloadOptions `json:"options,omitempty"`
}

type PayloadOptions struct {
	InterestedIn string `json:"interested_in,omitempty"`
	PhotoURL     string `json:"photo_url,omitempty"`
	Name         string `json:"name,omitempty"`
	Age          int    `json:"age,omitempty"`
	City         string `json:"city,omitempty"`
}

func NewPayload(command ButtonCommand) *Payload {
	return &Payload{
		Command: command,
	}
}

func UnmarshalPayload(rawPayload string) (payload Payload, err error) {
	err = json.Unmarshal([]byte(rawPayload), &payload)
	return
}
