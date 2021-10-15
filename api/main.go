package main

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
		Riddles.Riddle{Username: "Brendon", Question: "Make a fist", Answer: "And punch yourself"},
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
