package lrucache

import "github.com/gilwong00/go-lru-cache/pkg/queue"

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

func (l *lruCache) Add(n *queue.Node) error {
	return nil
}
