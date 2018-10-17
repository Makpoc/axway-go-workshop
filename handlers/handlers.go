package handlers

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
)

type postMessage struct {
	URL string `json:"url"`
}

type response struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}

func Shorten(w http.ResponseWriter, r *http.Request) {
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

	short := base64.URLEncoding.EncodeToString([]byte(msg.URL))

	responseBody := response{
		OriginalURL: msg.URL,
		ShortURL:    short,
	}

	// notice = instead of :=
	err = json.NewEncoder(w).Encode(responseBody)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}
