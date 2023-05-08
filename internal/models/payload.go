package models

import "encoding/json"

type ButtonCommand string

const (
	NoCommand           ButtonCommand = "no"
	StartCommand        ButtonCommand = "start"
	SexCommand          ButtonCommand = "sex"
	CreateCommand       ButtonCommand = "create"
	ShowCommand         ButtonCommand = "show"
	NameCommand         ButtonCommand = "name"
	CityCommand         ButtonCommand = "city"
	AgeCommand          ButtonCommand = "age"
	InterestedInCommand ButtonCommand = "interested_in"
	NextCommand         ButtonCommand = "next"
	LikeCommand         ButtonCommand = "like"
	DislikeCommand      ButtonCommand = "dislike"
	ReciprocityCommand  ButtonCommand = "reciprocity"
	GitHubCommand       ButtonCommand = "github"
	WhyISeeItCommand    ButtonCommand = "why"
)

type Payload struct {
	Command *ButtonCommand  `json:"command,omitempty"`
	Options *PayloadOptions `json:"options,omitempty"`
}

type PayloadOptions struct {
	InterestedIn string `json:"interested_in,omitempty"`
	PhotoURL     string `json:"photo_url,omitempty"`
	Name         string `json:"name,omitempty"`
	Age          int    `json:"age,omitempty"`
	City         string `json:"city,omitempty"`
	NoSave       *bool  `json:"no_save,omitempty"`
	ShownUserID  int    `json:"shown_user_id,omitempty"`
}

func NewPayload(cmd ButtonCommand, opts PayloadOptions) *Payload {
	return &Payload{
		Command: &cmd,
		Options: &opts,
	}
}

func NewPayloadWithCommandOnly(cmd ButtonCommand) *Payload {
	return &Payload{
		Command: &cmd,
	}
}

func UnmarshalPayload(rawPayload string) (payload Payload, err error) {
	err = json.Unmarshal([]byte(rawPayload), &payload)
	return
}
