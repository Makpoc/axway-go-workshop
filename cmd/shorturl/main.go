package main

import (
	"log"
	"net/http"
	"os"

	"github.com/makpoc/axway-go-workshop/handlers"
)

func main () {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on %s", port)

	http.HandleFunc("/", handlers.Welcome)
	http.HandleFunc("/postParser", handlers.PostParser)

	log.Fatal(http.ListenAndServe(":" + port, nil))
}
