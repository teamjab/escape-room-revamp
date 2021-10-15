package main

// Riddles obtained from https://parade.com/947956/parade/riddles/

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	Riddles "github.com/teamjab/escape-room-revamp/models"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"proof of life": "true"}`))
}

func getRiddle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
	json.NewEncoder(w).Encode(riddles)
}

func requestHandler() {
	router := mux.NewRouter()
	router.HandleFunc("/", homepage)
	router.HandleFunc("/riddles", getRiddle)

	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST"})

	log.Info("Server running on 8080")
	log.Error(http.ListenAndServe(":8080", handlers.CORS(originsOk, methodsOk)(router)))
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	requestHandler()
}
