package connectionhelper

import (
	"context"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client

var clientInstanceError error

var mongoOnce sync.Once

/*
Constant variables
*/
const (
	DBNAME       = "escape_room_revamp" // name of the db
	RIDDLE_ISSUE = "escape_riddle"
	SCORE_ISSUE = "escape_score"
)

/*
Connecting to the database once
*/
func ConnectDB() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}
