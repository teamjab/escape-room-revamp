package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"proof of life": "true"}`))
}

func requestHandler() {
	router := mux.NewRouter()
	router.HandleFunc("/", homepage)

	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, methodsOk)(router)))
}

func main() {
	requestHandler()
}
