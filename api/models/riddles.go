package Models

type Answers struct {
	AnswerOne   string
	AnswerTwo   string
	AnswerThree string
	AnswerFour  string
}

type Riddle struct {
	Username string    `bson:"Username"`
	Question string    `bson:"Question"`
	Answer   []Answers `bson:"Answers"`
	CorrectAnswer string `bson:"CorrectAnswer"`
}