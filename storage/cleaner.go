package storage

import (
	"log"
	"time"
)

func Clean(store Storage) {
	for {
		log.Println("Cleaning up")

		items := store.List()
		now := time.Now()
		for _, item := range items {
			if now.After(item.CreatedAt.Add(item.ExpireAfter)) {
				err := store.Delete(item.ShortID)
				if err != nil {
					log.Printf("Failed to clean up item with shortID: %s", item.ShortID)
				}
				log.Printf("Removed item with shortID: %s", item.ShortID)
			}
		}

		time.Sleep(5 * time.Second)
	}

}
