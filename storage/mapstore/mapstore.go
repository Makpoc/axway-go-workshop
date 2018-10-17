package mapstore

import (
	"sync"

	"github.com/makpoc/axway-go-workshop/storage"
)

// MapStore implements the storage interface with an underlying map as a key-value store.
type MapStore struct {
	// mux protects store
	mux   sync.Mutex
	store map[string]storage.Item
}

// New creates a new map store.
func New() *MapStore {
	return &MapStore{
		store: make(map[string]storage.Item),
	}
}

// Save stores the shortid and item pair in the map store. It returns storage.ShortIDAlreadyExistsErr if the shortid
// exists.
func (m *MapStore) Save(item storage.Item) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	if _, exists := m.store[item.ShortID]; exists {
		return storage.ShortIDAlreadyExistsErr
	}
	m.store[item.ShortID] = item
	// log.Printf("Saving [%s]=%s. Total entries: %d", shortid, url, len(m))
	return nil
}

// Load loads the item for given shortid from the map store. It returns storage.ShortIDNotFoundErr if the shortid cannot
// be found.
func (m *MapStore) Load(shortid string) (storage.Item, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	// _ means we will not be using the value
	if _, exists := m.store[shortid]; !exists {
		return storage.Item{}, storage.ShortIDNotFoundErr
	}
	return m.store[shortid], nil
}

// List returns a list with all the items in the current store
func (m *MapStore) List() []storage.Item {
	m.mux.Lock()
	defer m.mux.Unlock()

	v := make([]storage.Item, 0, len(m.store))

	for _, value := range m.store {
		v = append(v, value)
	}
	return v
}

// Delete deletes an item that matches the provided shortID from the store. It does nothing if the item doesn't exist
func (m *MapStore) Delete(shortID string) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	delete(m.store, shortID)
	return nil
}
