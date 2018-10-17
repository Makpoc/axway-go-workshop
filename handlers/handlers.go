// Package handlers contains all handlers for the shorturl service
package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/teris-io/shortid"
)

type postMessage struct {
	URL string `json:"url"`
}

type response struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}

// Handler holds the handleFunctions for the shorturl service
type Handler struct {
	// BaseURL is the schema://host:port/ of the service. It is used when generating the short URL
	BaseURL string
}

// New constructs a new Handler and configures it with the provided baseURL
func New(baseURL string) Handler {
	return Handler{baseURL}
}

// Shorten is a handleFunc that expects a POST request with a json payload and returns a response, containing the
// generated short URL
func (h Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Method %s not allowed for encoding", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var msg postMessage

	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		log.Printf("Failed to decode message: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	short, err := shortid.Generate()
	if err != nil {
		log.Printf("Failed to generate a short id: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody := response{
		OriginalURL: msg.URL,
		ShortURL:    h.BaseURL + short,
	}

	// notice = instead of :=
	err = json.NewEncoder(w).Encode(responseBody)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}
