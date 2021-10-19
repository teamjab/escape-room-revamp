package main

// Riddles obtained from https://parade.com/947956/parade/riddles/

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	RiddleController "github.com/teamjab/escape-room-revamp/controllers"
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
	riddles, err := RiddleController.GetRiddles()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(riddles)
}

func createRiddle(w http.ResponseWriter, r *http.Request) {
	var newRiddle Riddles.Riddle

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Error(err)
	}

	json.Unmarshal(reqBody, &newRiddle)

	status_code := RiddleController.CreateRiddles(newRiddle)

	if status_code == 200 {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content Type", "application/json")
		json.NewEncoder(w).Encode(newRiddle)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func requestHandler() {
	router := mux.NewRouter()
	router.HandleFunc("/", homepage).Methods("GET")
	router.HandleFunc("/riddles", getRiddle).Methods("GET")
	router.HandleFunc("/riddles", createRiddle).Methods("POST")

	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST"})

	log.Info("Server running on 8080")
	log.Error(http.ListenAndServe(":8080", handlers.CORS(originsOk, methodsOk)(router)))
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	godotenv.Load()
	requestHandler()
}
