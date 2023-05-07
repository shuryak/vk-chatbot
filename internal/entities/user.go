package entities

type User struct {
	VKID         int
	PhotoID      string
	Name         string
	Age          int
	City         string
	InterestedIn string
	Activated    *bool
}
