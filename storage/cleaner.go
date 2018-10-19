package storage

import (
	"log"
	"time"
)

// Cleaner performs storage maintenance on expired items.
type Cleaner struct {
	// Storage is the storage we will be maintaining
	Storage Storage
	// CleanInterval defines how often we will do the cleaning
	CleanInterval time.Duration
	// StopChan is the channel we will listen on for cleaning cancellation
	StopChan chan bool
}

// NewCleaner builds a new storage cleaner and returns a pointer to it
func NewCleaner(storage Storage, cleanInterval time.Duration, stopChan chan bool) *Cleaner {
	return &Cleaner{storage, cleanInterval, stopChan}
}

// Clean cleans up expired links.
func (c *Cleaner) Clean() {

CleanLoop:
	for {
		select {
		case <-time.Tick(c.CleanInterval):
			c.doClean()
		case <-c.StopChan:
			// Alternatively we can just return
			// log.Println("Cleaner stopped!")
			// return
			break CleanLoop
		}
	}
	log.Println("Cleaner stopped!")
}

// Stop stops the cleaner cycles
func (c *Cleaner) Stop() {
	c.StopChan <- true
}

// doClean performs the actual storage cleaning
func (c *Cleaner) doClean() {
	log.Println("Cleaning up")

	items := c.Storage.List()
	now := time.Now()
	for _, item := range items {
		if now.After(item.CreatedAt.Add(item.ExpireAfter)) {
			err := c.Storage.Delete(item.ShortID)
			if err != nil {
				log.Printf("Failed to clean up item with shortID: %s", item.ShortID)
			}
			log.Printf("Removed item with shortID: %s", item.ShortID)
		}
	}
}
