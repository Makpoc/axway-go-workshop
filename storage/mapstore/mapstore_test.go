package mapstore_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/makpoc/axway-go-workshop/storage"
	"github.com/makpoc/axway-go-workshop/storage/mapstore"
)

func TestNew(t *testing.T) {
	t.Parallel()

	store := mapstore.New()

	if store == nil {
		t.Fatalf("New() returned nil")
	}
}

func buildStorageItem(id int) storage.Item {
	// please ignore. I'm in a hurry :)
	return storage.Item{
		ShortID:     fmt.Sprintf("id%d", id),
		OriginalURL: fmt.Sprintf("url%d", id),
		CreatedAt:   time.Unix(0, 0),
		ExpireAfter: 1 * time.Second,
	}
}

func TestMapStore_Save(t *testing.T) {
	t.Parallel()

	store := mapstore.New()

	if len(store.List()) != 0 {
		t.Fatalf("Save(): expected len(store) 0, got %d", len(store.List()))
	}

	err := store.Save(buildStorageItem(1))
	if err != nil {
		t.Fatalf("Save(): expected no error, got %v", err)
	}
	if len(store.List()) != 1 {
		t.Fatalf("Save(): expected len(store) 1, got %d", len(store.List()))
	}

	err = store.Save(buildStorageItem(2))
	if err != nil {
		t.Fatalf("Save(): expected no error, got %v", err)
	}
	if len(store.List()) != 2 {
		t.Fatalf("Save(): expected len(store) 2, got %d", len(store.List()))
	}
}

func TestMapStore_Save_conflict(t *testing.T) {
	t.Parallel()

	store := mapstore.New()

	testItem := buildStorageItem(1)
	store.Save(testItem)
	err := store.Save(testItem)
	if err != storage.ShortIDAlreadyExistsErr {
		t.Fatalf("Save(): expected ShortIDAlreadyExistsErr, got %v", err)
	}
	if len(store.List()) != 1 {
		t.Fatalf("Save(): expected len(store) 1, got %d", len(store.List()))
	}
}

func TestMapStore_Load(t *testing.T) {
	t.Parallel()

	store := mapstore.New()

	store.Save(buildStorageItem(1))
	store.Save(buildStorageItem(2))
	store.Save(buildStorageItem(3))

	tt := []struct {
		testName     string
		givenShortId string
		expectErr    error
		expectItem   storage.Item
	}{
		{
			testName:     "first",
			givenShortId: "id1", // 😭
			expectErr:    nil,
			expectItem:   buildStorageItem(1),
		}, {
			testName:     "second",
			givenShortId: "id2", // 😭
			expectErr:    nil,
			expectItem:   buildStorageItem(2),
		}, {
			testName:     "third",
			givenShortId: "id3", // 😭
			expectErr:    nil,
			expectItem:   buildStorageItem(3),
		}, {
			testName:     "nonExistent",
			givenShortId: "id4",
			expectErr:    storage.ShortIDNotFoundErr,
			expectItem:   storage.Item{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(subTest *testing.T) {
			item, err := store.Load(tc.givenShortId)
			if err != tc.expectErr {
				subTest.Fatalf("Load(): expected no error, got %v", tc.expectErr)
			}
			if !reflect.DeepEqual(item, tc.expectItem) {
				subTest.Fatalf("Load(): expected %v, got %v", item, tc.expectItem)
			}
		})
	}
}

func ExampleMapStore() {
	store := mapstore.New()
	store.Save(storage.Item{ShortID: "id1", OriginalURL: "https://chucknorrisfacts.net/random-fact"})
	store.Save(storage.Item{ShortID: "id2", OriginalURL: "https://www.google.com/search?q=do+a+barrel+roll"})

	var item storage.Item
	item, _ = store.Load("id1")
	fmt.Println(item.OriginalURL)

	item, _ = store.Load("id2")
	fmt.Println(item.OriginalURL)

	// Output:
	// https://chucknorrisfacts.net/random-fact
	// https://www.google.com/search?q=do+a+barrel+roll
}

func BenchmarkMapStore_Save(b *testing.B) {
	store := mapstore.New()
	for n := 0; n < b.N; n++ {
		store.Save(storage.Item{
			ShortID:     fmt.Sprintf("id%d", n),
			OriginalURL: "https://chucknorrisfacts.net/random-fact"},
		)
	}
}
