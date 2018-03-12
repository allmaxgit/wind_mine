package store

import (
	"fmt"
	"os"
	"encoding/gob"

	"WindToken/services/BTCService/configs"

	"github.com/patrickmn/go-cache"
)

var store *cache.Cache

// Initiate store.
func InitiateStore() {
	fmt.Print("Initiate store...")

	items, err := getFromFile()
	if err != nil {
		fmt.Println("failed to rise store from file")
	}

	if len(items) > 0 {
		store = cache.NewFrom(cache.NoExpiration, cache.NoExpiration, items)
	} else {
		store = cache.New(cache.NoExpiration, cache.NoExpiration)
	}

	return
}

// Setter for cache.
// Sets value without expiration.
func Set(key string, value interface{}) {
	store.Set(key, value, cache.NoExpiration)
}

// Getter for cache.
func Get(key string) (value interface{}, found bool) {
	value, found = store.Get(key)
	return
}

// Save saves store items to file.
func Save() (err error) {
	conf := configs.GetConfigs()
	items := store.Items()

	f, err := os.OpenFile(conf.Common.StorePath, os.O_WRONLY, os.ModeAppend)
	if err != nil { return }
	defer f.Close()

	encoder := gob.NewEncoder(f)
	err = encoder.Encode(items)

	return
}

func getFromFile() (map[string]cache.Item, error) {
	conf := configs.GetConfigs()
	var items map[string]cache.Item

	f, err := os.OpenFile(conf.Common.StorePath, os.O_RDONLY, os.ModeAppend)
	if err != nil { return items, err }
	defer f.Close()

	decoder := gob.NewDecoder(f)
	err = decoder.Decode(&items)

	return items, err
}
