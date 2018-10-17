package main

import (
	"log"
	"net/http"
)

func main () {
	port := "6789"
	log.Printf("Starting server on %s", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
