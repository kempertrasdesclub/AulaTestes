package main

import (
	"fmt"
	"github.com/dgraph-io/ristretto"
)

func main() {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}

	// set a value with a cost of 1
	cache.Set("key", "value", 1)

	// wait for value to pass through buffers
	cache.Wait()

	value, found := cache.Get("key")
	if !found {
		panic("missing value")
	}
	fmt.Println(value)
	cache.Del("key")
}
