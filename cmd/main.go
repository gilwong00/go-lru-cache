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
	lruCache.Add(&queue.Node{Val: "test"})
	// print cache values
	lruCache.Print()
	// Remove

	// print
}
