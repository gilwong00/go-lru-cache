package lrucache

import (
	"fmt"

	"github.com/gilwong00/go-lru-cache/pkg/queue"
)

type lruCache struct {
	Queue *queue.Queue
	Hash  queue.Hash
	size  int64
}

func newLruCache(size int64) *lruCache {
	return &lruCache{
		size:  size,
		Queue: queue.NewQueue(),
		Hash:  queue.Hash{},
	}
}

func (l *lruCache) Add(n *queue.Node) {
	fmt.Println("adding: %s", n.Val)
	// need to check if the incoming record already exist, if does, remove and readd
	copy := l.Queue.Head.Right
	l.Queue.Head.Right = n
	n.Left = l.Queue.Head
	n.Right = copy
	copy.Left = n
	l.Queue.Length++
	// TODO: check if we hit size limit
	// if l.Queue.Length > int(l.size) {

	// }
}

func (l *lruCache) Print() {
	node := l.Queue.Head.Right
	fmt.Printf("%d - [", l.Queue.Length)
	for i := 0; i < l.Queue.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < l.Queue.Length-1 {
			fmt.Printf("<->")
		}
		node = node.Right
	}
	fmt.Println("]")
}
