package storage

import (
	"fmt"
	"time"
)

// Storage describes the interface any implementation must comply to
type Storage interface {
	// Save saves an item in the storage.
	Save(item Item) error
	// Load loads the item for given shortid. It returns an error if something goes wrong. It will return the
	// concrete error ShortIDNotFoundErr if we have no record for the provided shortid
	Load(shortID string) (Item, error)
	// List returns all items in the store
	List() []Item
	// Delete deletes an item from the storage
	Delete(shortID string) error
}

// ShortIDNotFoundErr is the error returned when an item matching the shortid was not found
var ShortIDNotFoundErr = fmt.Errorf("shortid not found")

// ShortIDAlreadyExistsErr is the error returned when trying to save an item with already existing shortid.
var ShortIDAlreadyExistsErr = fmt.Errorf("conflict: shortid already exists")

// Item is the struct, stored in a storage.
type Item struct {
	// ShortID is the generated short ID for the item
	ShortID string
	// OriginalURL is the URL mapped to the shortId.
	OriginalURL string
	// CreatedAt is the datetime the shortId was created
	CreatedAt   time.Time
	// ExpireAfter is the time the item will expire and will be removed from the store
	ExpireAfter time.Duration
}
