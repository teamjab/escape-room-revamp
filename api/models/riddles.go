package Riddles

type Answers struct {
	AnswerOne   string
	AnswerTwo   string
	AnswerThree string
	AnswerFour  string
}

type Riddle struct {
	Username string    `json:"Username"`
	Question string    `json:"Question"`
	Answer   []Answers `json:"Answers"`
}
