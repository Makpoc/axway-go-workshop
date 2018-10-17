package mapstore

import (
	"log"

	"github.com/makpoc/axway-go-workshop/storage"
)

// MapStore implements the storage interface with an underlying map as a key-value store.
type MapStore map[string]string

// New creates a new map store.
func New() MapStore {
	var store = make(map[string]string)
	return MapStore(store)
}

// Save stores the shortid and url pair in the map store. It returns storage.ShortIDAlreadyExistsErr if the shortid
// exists.
func (m MapStore) Save(shortid, url string) error {
	// _ means we will not be using the value
	if _, exists := m[shortid]; exists {
		return storage.ShortIDAlreadyExistsErr
	}
	m[shortid] = url
	log.Printf("Saving [%s]=%s. Total entries: %d", shortid, url, len(m))
	return nil
}

// Load loads the url for given shortid from the map store. It returns storage.ShortIDNotFoundErr if the shortid cannot
// be found.
func (m MapStore) Load(shortid string) (string, error) {
	// _ means we will not be using the value
	if _, exists := m[shortid]; !exists {
		return "", storage.ShortIDNotFoundErr
	}
	return m[shortid], nil
}
