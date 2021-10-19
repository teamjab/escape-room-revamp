package ApiHandlers

import (
	"context"

	log "github.com/sirupsen/logrus"
	Models "github.com/teamjab/escape-room-revamp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	connectionhelper "github.com/teamjab/escape-room-revamp/db"
)

/*
Saving the riddles that are passed from the front end to the db
*/
func CreateRiddles(riddle Models.Riddle) int {
	if riddle.Answer == nil || riddle.Question == "" || riddle.Username == "" {
		return 400
	}

	client, err := connectionhelper.ConnectDB()
	if err != nil {
		log.Error(err)
		return 400
	}

	log.Info("Posting riddle to the db")

	log.Debug("This is riddle that are being stored {}", riddle)

	collection := client.Database(connectionhelper.DBNAME).Collection(connectionhelper.RIDDLE_ISSUE)

	_, err = collection.InsertOne(context.TODO(), riddle)

	if err != nil {
		log.Error(err)
		return 400
	} else {
		return 200
	}
}

/*
Getting all of the riddles from the db and mock riddles
*/
func GetRiddles() ([]Models.Riddle, error) {
	log.Info("Getting riddles from the db...")

	riddles := mockRiddles()

	filter := bson.D{{}}

	client, err := connectionhelper.ConnectDB()

	if err != nil {
		return riddles, err
	}

	collection := client.Database(connectionhelper.DBNAME).Collection(connectionhelper.RIDDLE_ISSUE)

	riddlesCol, dbError := collection.Find(context.TODO(), filter)

	if dbError != nil {
		return riddles, dbError
	}

	for riddlesCol.Next(context.TODO()) {
		rid := Models.Riddle{}
		err := riddlesCol.Decode(&rid)
		if err != nil {
			return riddles, err
		}
		riddles = append(riddles, rid)
	}

	riddlesCol.Close(context.TODO())
	if len(riddles) == 0 {
		return riddles, mongo.ErrNoDocuments
	}
	return riddles, nil
}

/*
Function to create mock riddles
*/
func mockRiddles() []Models.Riddle {
	return []Models.Riddle{
		{
			Username: "Brendon",
			Question: "Make a fist",
			Answer: []Models.Answers{
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
			Answer: []Models.Answers{
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
			Answer: []Models.Answers{
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
			Answer: []Models.Answers{
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
			Answer: []Models.Answers{
				{
					AnswerOne:   "My car",
					AnswerTwo:   "Air plane",
					AnswerThree: "Yo mama",
					AnswerFour:  "A stamp",
				},
			},
		},
	}
}
