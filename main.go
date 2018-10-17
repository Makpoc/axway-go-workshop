package main

import (
	"log"
	"net/http"
	"os"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome"))
}

func main () {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on %s", port)

	http.HandleFunc("/", welcome)

	log.Fatal(http.ListenAndServe(":" + port, nil))
}
