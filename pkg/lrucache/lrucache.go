package lrucache

import (
	"fmt"

	"github.com/gilwong00/go-lru-cache/pkg/queue"
)

type lruCache struct {
	Queue      *queue.Queue
	Dictionary queue.Dictionary
	size       int64
}

func newLruCache(size int64) *lruCache {
	return &lruCache{
		size:       size,
		Queue:      queue.NewQueue(),
		Dictionary: queue.Dictionary{},
	}
}

func (l *lruCache) Add(n *queue.Node) {
	fmt.Printf("adding: %s\n", n.Val)
	// check if the incoming node already exist, if does, remove and readd to HEAD
	if _, ok := l.Dictionary[n.Val]; ok {
		l.Delete(n)
	}
	copy := l.Queue.Head.Right
	l.Queue.Head.Right = n
	n.Left = l.Queue.Head
	n.Right = copy
	copy.Left = n
	l.Queue.Length++
	l.Dictionary[n.Val] = n
	// if we hit size limit, remove the tail
	if l.Queue.Length > int(l.size) {
		l.Delete(l.Queue.Tail.Left)
	}
}

func (l *lruCache) Delete(n *queue.Node) {
	fmt.Printf("removing: %s from cache\n", n.Val)
	left := n.Left
	right := n.Right
	if right != nil {
		left.Right = right
	}
	if left != nil {
		left.Left = left
	}
	l.Queue.Length -= 1
	delete(l.Dictionary, n.Val)
}

func (l *lruCache) Print() {
	node := l.Queue.Head.Right
	fmt.Printf("%d - [ ", l.Queue.Length)
	for i := 0; i < l.Queue.Length; i++ {
		fmt.Printf("(%s)", node.Val)
		if i < l.Queue.Length-1 {
			fmt.Printf(" <---> ")
		}
		node = node.Right
	}
	fmt.Println(" ]")
}
