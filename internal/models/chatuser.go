package models

import "time"

type ChatUser struct {
	ID        int
	PhotoID   string
	Name      string
	BirthDate *time.Time
	City      string
}
