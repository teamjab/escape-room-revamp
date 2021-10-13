package main

import (
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"proof of life": "true"}`))
}

func requestHandler() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	requestHandler()
}
