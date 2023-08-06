package main

import (
	"fmt"

	"github.com/gilwong00/go-lru-cache/pkg/lrucache"
	"github.com/gilwong00/go-lru-cache/pkg/queue"
)

func main() {
	fmt.Println("CACHE INITIALIZE")
	lruCache := lrucache.NewLruCache(10)

	// Add
	err := lruCache.Add(&queue.Node{Val: "test"})
	if err != nil {
		fmt.Printf("error adding to cache: %v", err)
		panic(err)
	}

	// print cache values

	// Remove

	// print
}
