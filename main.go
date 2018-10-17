package main

import (
	"log"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome"))
}

func main () {
	port := "6789"
	log.Printf("Starting server on %s", port)

	http.HandleFunc("/", welcome)

	log.Fatal(http.ListenAndServe(":" + port, nil))
}
