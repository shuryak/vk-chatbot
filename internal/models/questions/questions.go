package questions

type QuestionType string

const (
	NoQuestion    QuestionType = "no"
	PhotoQuestion QuestionType = "photo"
	AgeQuestion   QuestionType = "age"
	CityQuestion  QuestionType = "city"
	NameQuestion  QuestionType = "name"
)
