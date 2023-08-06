package main

import (
	"fmt"

	"github.com/gilwong00/go-lru-cache/pkg/lrucache"
	"github.com/gilwong00/go-lru-cache/pkg/queue"
)

func main() {
	fmt.Println("CACHE INITIALIZING")
	lruCache := lrucache.NewLruCache(5)
	// Add
	lruCache.Add(&queue.Node{Val: "test"})
	lruCache.Add(&queue.Node{Val: "meep"})
	// print cache values
	lruCache.Print()
	// Delete
	lruCache.Delete(&queue.Node{Val: "test"})
	// print
	lruCache.Print()
	// pop from tail if limit is hit
	lruCache.Add(&queue.Node{Val: "test"})
	lruCache.Add(&queue.Node{Val: "meep"})
	lruCache.Add(&queue.Node{Val: "morp"})
	lruCache.Add(&queue.Node{Val: "foo"})
	lruCache.Add(&queue.Node{Val: "bar"})
	lruCache.Print()
	lruCache.Add(&queue.Node{Val: "yeet"})
	lruCache.Print()
}
