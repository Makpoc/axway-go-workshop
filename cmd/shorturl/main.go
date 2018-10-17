package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/makpoc/axway-go-workshop/storage"
	"github.com/makpoc/axway-go-workshop/storage/mapstore"
	"github.com/teris-io/shortid"

	"github.com/makpoc/axway-go-workshop/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "6789"
	}
	log.Printf("Starting server on %s", port)

	// create and configure a shortid generator
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		log.Fatalf("Failed to create shortid generator: %v", err)
	}
	shortid.SetDefault(sid)

	baseURL := fmt.Sprintf("http://localhost:%s/redirect/", port)
	store := mapstore.New()
	handler := handlers.New(baseURL, store)

	// Spawns a new goroutine
	go storage.Clean(store)

	http.HandleFunc("/shorten", handler.Shorten)
	// note the trailing slash - this means match /redirect/*
	http.HandleFunc("/redirect/", handler.Redirect)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
