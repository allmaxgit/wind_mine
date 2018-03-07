package store

import (
	"fmt"

	"github.com/patrickmn/go-cache"
)

var Store *cache.Cache

// Initiate store
func InitiateStore() {
	fmt.Print("Initiate store...")
	Store = cache.New(cache.NoExpiration, cache.NoExpiration)
}

// Setter for cache.
// Sets value without expiration
func Set(key string, value interface{}) {
	Store.Set(key, value, cache.NoExpiration)
}

// Getter for cache.
func Get(key string) (value interface{}, found bool) {
	value, found = Store.Get(key)
	return
}
