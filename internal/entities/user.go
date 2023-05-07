package entities

type User struct {
	VKID         int
	PhotoURL     string
	Name         string
	Age          int
	City         string
	InterestedIn string
	Activated    *bool
}
