package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"proof of life": "true"}`))
}

func requestHandler() {
	router := mux.NewRouter()
	router.HandleFunc("/", homepage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	requestHandler()
}
