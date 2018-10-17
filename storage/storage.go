package storage

import "fmt"

// Storage describes the interface any implementation must comply to
type Storage interface {
	// Save saves a shortid to real url mapping. It returns an error if saving fails
	Save(shortid, url string) error
	// Load loads the real url for given shortid. It returns an error if something goes wrong. It will return the
	// concrete error ShortIDNotFoundErr if we have no record for the provided shortid
	Load(shortid string) (string, error)
}

// ShortIDNotFoundErr is the error returned when an item matching the shortid was not found
var ShortIDNotFoundErr = fmt.Errorf("shortid not found")

// ShortIDAlreadyExistsErr is the error returned when trying to save an item with already existing shortid.
var ShortIDAlreadyExistsErr = fmt.Errorf("conflict: shortid already exists")
