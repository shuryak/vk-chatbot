package models

type Sympathy struct {
	ID           int
	FirstUserID  int
	SecondUserID int
	Reciprocity  bool
}
