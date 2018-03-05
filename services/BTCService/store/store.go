package store

import (
	"fmt"

	"github.com/patrickmn/go-cache"
)

var store *cache.Cache

func InitiateStore() {
	fmt.Print("Initiate store")
	store = cache.New(cache.NoExpiration, cache.NoExpiration)
}
