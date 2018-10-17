package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome"))
}

type postMessage struct {
	URL string `json:"url"`
}

func postParser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Method %s not allowed for encoding", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var msg postMessage

	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		log.Printf("Failed to decode message: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Got %v", msg)
}
