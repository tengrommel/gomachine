package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {
	// Create a cache with a default expiration time of 5 minutes,
	// and which purges expired items every
	c := cache.New(5*time.Minute, 30*time.Second)
	// Put a key and value into the cache
	c.Set("mykey", "myvalue", cache.DefaultExpiration)
	// For a sanity check. Output the key and value in the cache to standard out.
	v, found := c.Get("mykey")
	if found {
		fmt.Printf("key: mykey, value: %s\n", v)
	}
}
