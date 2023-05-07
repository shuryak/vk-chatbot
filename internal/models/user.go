package models

type User struct {
	ID           int
	PhotoID      string
	Name         string
	Age          int
	City         string
	InterestedIn string
	Activated    *bool
}
