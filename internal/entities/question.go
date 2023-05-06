package entities

type QuestionType string

const (
	NoQuestion    QuestionType = "no"
	PhotoQuestion QuestionType = "photo"
	AgeQuestion   QuestionType = "age"
	CityQuestion  QuestionType = "city"
	NameQuestion  QuestionType = "name"
)

func (qt QuestionType) MarshalBinary() ([]byte, error) {
	return []byte(qt), nil
}

func (qt *QuestionType) UnmarshalBinary(data []byte) error {
	*qt = QuestionType(data)
	return nil
}
