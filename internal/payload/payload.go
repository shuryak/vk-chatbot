package payload

import "encoding/json"

type ButtonPayload string

const (
	Start  ButtonPayload = "start"
	Create ButtonPayload = "create"
)

type Payload struct {
	Command ButtonPayload `json:"command"`
}

func New(command ButtonPayload) *Payload {
	return &Payload{command}
}

func Unmarshal(rawPayload string) (payload Payload, err error) {
	err = json.Unmarshal([]byte(rawPayload), &payload)
	return
}
