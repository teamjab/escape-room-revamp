package ApiHandlers

import (
	"context"

	log "github.com/sirupsen/logrus"
	Models "github.com/teamjab/escape-room-revamp/models"

	connectionhelper "github.com/teamjab/escape-room-revamp/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
Saving the score to the db
Score has to be int type and username has to be string type
*/
func CreateScore(scoreInfo Models.Scores) int {
	if scoreInfo.Username == "" {
		log.Error("Invalid input...")
		return 400
	}

	client, err := connectionhelper.ConnectDB()
	if err != nil {
		log.Error(err)
		return 400
	}

	log.Info("Posting score to the db")

	collection := client.Database(connectionhelper.DBNAME).Collection(connectionhelper.SCORE_ISSUE)

	_, err = collection.InsertOne(context.TODO(), scoreInfo)

	if err != nil {
		log.Error(err)
		return 400
	} else {
		return 200
	}
}

/*
Controller for getting all of the scores
*/
func GetAllScore() ([]Models.Scores, error) {
	log.Info("Getting all the scores from the db...")

	filter := bson.D{{}}

	aggregateScores := []Models.Scores{}

	client, err := connectionhelper.ConnectDB()

	if err != nil {
		return aggregateScores, err
	}

	collection := client.Database(connectionhelper.DBNAME).Collection(connectionhelper.SCORE_ISSUE)

	scoresCol, dbError := collection.Find(context.TODO(), filter)

	if dbError != nil {
		return aggregateScores, dbError
	}

	for scoresCol.Next(context.TODO()) {
		score := Models.Scores{}
		err := scoresCol.Decode(&score)
		if err != nil {
			return aggregateScores, err
		}
		aggregateScores = append(aggregateScores, score)
	}

	scoresCol.Close(context.TODO())
	if len(aggregateScores) == 0 {
		return aggregateScores, mongo.ErrNoDocuments
	}
	return aggregateScores, nil
}
