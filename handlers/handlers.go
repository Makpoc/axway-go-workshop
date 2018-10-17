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

type Handler struct {
	BaseURL string
}

func New(baseURL string) Handler {
	return Handler{baseURL}
}

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
