package mapstore_test

import (
	"fmt"
	"testing"

	"github.com/makpoc/axway-go-workshop/storage"
	"github.com/makpoc/axway-go-workshop/storage/mapstore"
)

func TestNew(t *testing.T) {
	store := mapstore.New()

	if store == nil {
		t.Fatalf("New() returned nil")
	}
}

func TestMapStore_Save(t *testing.T) {
	store := mapstore.New()

	if len(store) != 0 {
		t.Fatalf("Save(): expected len(store) 0, got %d", len(store))
	}

	err := store.Save("id1", "url1")
	if err != nil {
		t.Fatalf("Save(): expected no error, got %v", err)
	}
	if len(store) != 1 {
		t.Fatalf("Save(): expected len(store) 1, got %d", len(store))
	}

	err = store.Save("id2", "url2")
	if err != nil {
		t.Fatalf("Save(): expected no error, got %v", err)
	}
	if len(store) != 2 {
		t.Fatalf("Save(): expected len(store) 2, got %d", len(store))
	}
}

func TestMapStore_Save_conflict(t *testing.T) {
	store := mapstore.New()

	store.Save("id1", "url1")
	err := store.Save("id1", "url1")
	if err != storage.ShortIDAlreadyExistsErr {
		t.Fatalf("Save(): expected ShortIDAlreadyExistsErr, got %v", err)
	}
	if len(store) != 1 {
		t.Fatalf("Save(): expected len(store) 1, got %d", len(store))
	}
}

func TestMapStore_Load(t *testing.T) {
	store := mapstore.New()

	store.Save("id1", "url1")
	store.Save("id2", "url2")
	store.Save("id3", "url3")

	tt := []struct {
		testName     string
		givenShortId string
		expectErr    error
		expectURL    string
	}{
		{
			testName:     "first",
			givenShortId: "id1",
			expectErr:    nil,
			expectURL:    "url1",
		}, {
			testName: "second",
			givenShortId: "id2",
			expectErr:    nil,
			expectURL:    "url2",
		}, {
			testName: "third",
			givenShortId: "id3",
			expectErr:    nil,
			expectURL:    "url3",
		}, {
			testName: "nonExistent",
			givenShortId: "id4",
			expectErr:    storage.ShortIDNotFoundErr,
			expectURL:    "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(subTest *testing.T) {
			url, err := store.Load(tc.givenShortId)
			if err != tc.expectErr {
				t.Fatalf("Load(): expected no error, got %v", tc.expectErr)
			}
			if url != tc.expectURL {
				t.Fatalf("Load(): expected url1, got %s", tc.expectURL)
			}
		})
	}
}

func ExampleMapStore() {
	store := mapstore.New()
	store.Save("id1", "https://chucknorrisfacts.net/random-fact")
	store.Save("id2", "https://www.google.com/search?q=do+a+barrel+roll")

	var url string
	url, _ = store.Load("id1")
	fmt.Println(url)

	url, _ = store.Load("id2")
	fmt.Println(url)
	// Output:
	// https://chucknorrisfacts.net/random-fact
	// https://www.google.com/search?q=do+a+barrel+roll
}
