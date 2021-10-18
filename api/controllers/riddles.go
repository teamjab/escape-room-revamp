package RiddleController

import (
	"context"

	log "github.com/sirupsen/logrus"
	Riddles "github.com/teamjab/escape-room-revamp/models"

	connectionhelper "github.com/teamjab/escape-room-revamp/db"
)

func CreateRiddles(riddle Riddles.Riddle) int {
	if riddle.Answer == nil || riddle.Question == "" || riddle.Username == "" {
		return 400
	}

	client, err := connectionhelper.ConnectDB()
	if err != nil {
		log.Error(err)
		return 400
	}

	collection := client.Database(connectionhelper.DBNAME).Collection(connectionhelper.RIDDLE_ISSUE)

	_, err = collection.InsertOne(context.TODO(), riddle)

	if err != nil {
		log.Error(err)
		return 400
	} else {
		return 200
	}
}

func GetRiddles() []Riddles.Riddle {
	var riddles = []Riddles.Riddle{
		{
			Username: "Brendon",
			Question: "Make a fist",
			Answer: []Riddles.Answers{
				{
					AnswerOne:   "Punch yourself",
					AnswerTwo:   "Punch Brendon",
					AnswerThree: "Punch Brendon",
					AnswerFour:  "Punch Brendon",
				},
			},
		},
		{
			Username: "TeamJab",
			Question: "What has to be broken before you can use it?",
			Answer: []Riddles.Answers{
				{
					AnswerOne:   "An egg",
					AnswerTwo:   "A life",
					AnswerThree: "Brendon's face",
					AnswerFour:  "Jin's face",
				},
			},
		},
		{
			Username: "TeamJab",
			Question: "What month of the year has 28 days?",
			Answer: []Riddles.Answers{
				{
					AnswerOne:   "Febuary",
					AnswerTwo:   "March",
					AnswerThree: "All of them",
					AnswerFour:  "DecemberMarch",
				},
			},
		},
		{
			Username: "TeamJab",
			Question: "What has a head and a tail but no body?",
			Answer: []Riddles.Answers{
				{
					AnswerOne:   "Brendon",
					AnswerTwo:   "Jin",
					AnswerThree: "Life",
					AnswerFour:  "Coin",
				},
			},
		},
		{
			Username: "TeamJab",
			Question: "What can travel all around the world without leaving its corner?",
			Answer: []Riddles.Answers{
				{
					AnswerOne:   "My car",
					AnswerTwo:   "Air plane",
					AnswerThree: "Yo mama",
					AnswerFour:  "A stamp",
				},
			},
		},
	}
	return riddles
}
