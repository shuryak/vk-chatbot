package models

import "time"

type User struct {
	ID        int
	PhotoID   string
	Name      string
	BirthDate *time.Time
	City      string
}
