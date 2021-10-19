package Models

type Scores struct {
	Username string `bson:"Username"`
	Score    int    `bson:"Score"`
}
