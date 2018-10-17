package mapstore

import (
	"github.com/makpoc/axway-go-workshop/storage"
)

// mapStore implements the storage interface with an underlying map as a key-value store.
type mapStore map[string]string

// New creates a new map store.
func New() mapStore {
	var store = make(map[string]string)
	return mapStore(store)
}

// Save stores the shortid and url pair in the map store. It returns storage.ShortIDAlreadyExistsError if the shortid
// exists.
func (m mapStore) Save(shortid, url string) error {
	// _ means we will not be using the value
	if _, exists := m[shortid]; exists {
		return storage.ShortIDAlreadyExistsError
	}
	m[shortid] = url
	return nil
}

// Load loads the url for given shortid from the map store. It returns storage.ShortIDNotFoundErr if the shortid cannot
// be found.
func (m mapStore) Load(shortid string) (string, error) {
	// _ means we will not be using the value
	if _, exists := m[shortid]; !exists {
		return "", storage.ShortIDNotFoundErr
	}
	return m[shortid], nil
}
