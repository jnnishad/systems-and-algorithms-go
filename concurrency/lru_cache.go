// Package concurrency holds the "systems" half of this repo: the kind
// of building blocks that show up in infra tooling (rate limiters,
// worker pools, caches) more often than pure LeetCode-style puzzles do,
// implemented with explicit attention to concurrency safety.
package concurrency

import (
	"container/list"
	"sync"
)

// LRUCache is a fixed-capacity, goroutine-safe least-recently-used
// cache. O(1) Get and Put via a doubly linked list (container/list) +
// a map from key to list element, protected by a single mutex.
type LRUCache struct {
	mu       sync.Mutex
	capacity int
	ll       *list.List
	items    map[string]*list.Element
}

type entry struct {
	key   string
	value any
}

func NewLRUCache(capacity int) *LRUCache {
	if capacity <= 0 {
		capacity = 1
	}
	return &LRUCache{
		capacity: capacity,
		ll:       list.New(),
		items:    make(map[string]*list.Element, capacity),
	}
}

// Get returns the value for key and marks it most-recently-used.
func (c *LRUCache) Get(key string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	el, ok := c.items[key]
	if !ok {
		return nil, false
	}
	c.ll.MoveToFront(el)
	return el.Value.(*entry).value, true
}

// Put inserts or updates key, evicting the least-recently-used entry
// if the cache is at capacity.
func (c *LRUCache) Put(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, ok := c.items[key]; ok {
		el.Value.(*entry).value = value
		c.ll.MoveToFront(el)
		return
	}

	if c.ll.Len() >= c.capacity {
		oldest := c.ll.Back()
		if oldest != nil {
			c.ll.Remove(oldest)
			delete(c.items, oldest.Value.(*entry).key)
		}
	}

	el := c.ll.PushFront(&entry{key: key, value: value})
	c.items[key] = el
}

func (c *LRUCache) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ll.Len()
}
