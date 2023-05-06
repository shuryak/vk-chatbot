package entities

type Sympathy struct {
	ID             int
	FirstUserVKID  int
	SecondUserVKID int
	Reciprocity    bool
}
