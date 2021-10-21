package main

// Riddles obtained from https://parade.com/947956/parade/riddles/

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	ApiHandlers "github.com/teamjab/escape-room-revamp/controllers"
	Models "github.com/teamjab/escape-room-revamp/models"
)

/*
Proof of life homepage
*/
func homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"proof of life": "true"}`))
}

/*
Getting all the riddles that are premade and from the DB
*/
func getRiddle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	riddles, err := ApiHandlers.GetRiddles()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(riddles)
}

/*
Creating a riddles and posting it to the DB
*/
func createRiddle(w http.ResponseWriter, r *http.Request) {
	var newRiddle Models.Riddle

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Error(err)
	}

	json.Unmarshal(reqBody, &newRiddle)

	status_code := ApiHandlers.CreateRiddles(newRiddle)

	if status_code == 200 {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content Type", "application/json")
		json.NewEncoder(w).Encode(newRiddle)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Bad Request")
	}
}

/*
Getting all the scores from the DB
*/
func getScores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	aggregateScores, err := ApiHandlers.GetAllScore()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(aggregateScores)
}

/*
Creating a score info and posting it to the DB
*/
func createScoreInfo(w http.ResponseWriter, r *http.Request) {
	var scoreInfo Models.Scores

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Error(err)
	}

	json.Unmarshal(reqBody, &scoreInfo)

	log.Info(scoreInfo)
	status_code := ApiHandlers.CreateScore(scoreInfo)

	if status_code == 200 {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content Type", "application/json")
		json.NewEncoder(w).Encode(scoreInfo)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Bad Request")
	}
}

/*
All of the routes
*/
func requestHandler() {
	router := mux.NewRouter()
	router.HandleFunc("/", homepage).Methods("GET")
	router.HandleFunc("/riddles", getRiddle).Methods("GET")
	router.HandleFunc("/riddles", createRiddle).Methods("POST")
	router.HandleFunc("/scores", createScoreInfo).Methods("POST")
	router.HandleFunc("/scores", getScores).Methods("GET")

	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST"})

	PORT := os.Getenv("PORT")
	log.Info("Server running on " + PORT)
	log.Error(http.ListenAndServe(":"+PORT, handlers.CORS(originsOk, methodsOk)(router)))
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	godotenv.Load()
	requestHandler()
}
