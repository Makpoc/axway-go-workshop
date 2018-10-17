// Package handlers contains all handlers for the shorturl service
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/makpoc/axway-go-workshop/storage"
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
	Storage storage.Storage
}

// New constructs a new Handler and configures it with the provided baseURL
func New(baseURL string, storage storage.Storage) Handler {
	return Handler{baseURL, storage}
}

// Shorten is a handleFunc that expects a POST request with a json payload and returns a response, containing the
// generated short URL
func (h Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Method %s not allowed for shorten", r.Method)
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

	err = h.Storage.Save(short, storage.Item{
		ShortID: short,
		OriginalURL: msg.URL,
		CreatedAt:   time.Now(),
		ExpireAfter: 10 * time.Second,
	})
	if err != nil {
		log.Printf("Failed to store the short id: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody := response{
		OriginalURL: msg.URL,
		ShortURL:    h.BaseURL + short,
	}

	err = json.NewEncoder(w).Encode(responseBody)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

// Redirect redirects the user to the original URL based on provided shortid.
func (h Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("Method %s not allowed for redirect", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	short := getCleanShortFromPath(r.URL.Path)
	if short == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	item, err := h.Storage.Load(short)
	if err != nil {
		log.Printf("Failed to load the short id: %v", err)
		if err == storage.ShortIDNotFoundErr {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Location", item.OriginalURL)
	w.WriteHeader(http.StatusMovedPermanently)
}

func getCleanShortFromPath(path string) string {
	short := strings.TrimPrefix(path, "/redirect")
	short = strings.TrimPrefix(short, "/")
	return strings.TrimSpace(short)
}
