package mapstore

import (
	"github.com/makpoc/axway-go-workshop/storage"
)

// MapStore implements the storage interface with an underlying map as a key-value store.
type MapStore map[string]storage.Item

// New creates a new map store.
func New() MapStore {
	var store = make(map[string]storage.Item)
	return MapStore(store)
}

// Save stores the shortid and item pair in the map store. It returns storage.ShortIDAlreadyExistsErr if the shortid
// exists.
func (m MapStore) Save(item storage.Item) error {
	// _ means we will not be using the value
	if _, exists := m[item.ShortID]; exists {
		return storage.ShortIDAlreadyExistsErr
	}
	m[item.ShortID] = item
	// log.Printf("Saving [%s]=%s. Total entries: %d", shortid, url, len(m))
	return nil
}

// Load loads the item for given shortid from the map store. It returns storage.ShortIDNotFoundErr if the shortid cannot
// be found.
func (m MapStore) Load(shortid string) (storage.Item, error) {
	// _ means we will not be using the value
	if _, exists := m[shortid]; !exists {
		return storage.Item{}, storage.ShortIDNotFoundErr
	}
	return m[shortid], nil
}

func (m MapStore) List() []storage.Item {
	v := make([]storage.Item, 0, len(m))
	for _, value := range m {
		v = append(v, value)
	}
	return v
}

func (m MapStore) Delete(shortID string) error {
	delete(m, shortID)
	return nil
}
