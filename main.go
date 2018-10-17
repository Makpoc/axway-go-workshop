package main

import (
	"log"
	"net/http"
	"os"
)

func main () {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on %s", port)

	http.HandleFunc("/", welcome)
	http.HandleFunc("/postParser", postParser)


	log.Fatal(http.ListenAndServe(":" + port, nil))
}
