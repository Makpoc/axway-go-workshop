package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/teris-io/shortid"

	"github.com/makpoc/axway-go-workshop/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on %s", port)

	// create and configure a shortid generator
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		log.Fatalf("Failed to create shortid generator: %v", err)
	}
	shortid.SetDefault(sid)

	handler := handlers.New(fmt.Sprintf("http://localhost:%s/", port))

	http.HandleFunc("/shorten", handler.Shorten)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}