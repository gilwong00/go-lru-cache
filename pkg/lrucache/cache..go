package lrucache

import "github.com/gilwong00/go-lru-cache/pkg/queue"

type LruCache interface {
	Add(n *queue.Node) error
}

func NewLruCache(size int64) LruCache {
	return newLruCache(size)
}
