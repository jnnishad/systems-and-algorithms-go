package concurrency

import (
	"sync"
	"testing"
)

func TestLRUCache_BasicGetPut(t *testing.T) {
	c := NewLRUCache(2)
	c.Put("a", 1)
	c.Put("b", 2)

	if v, ok := c.Get("a"); !ok || v != 1 {
		t.Fatalf("expected a=1, got %v, ok=%v", v, ok)
	}
}

func TestLRUCache_EvictsLeastRecentlyUsed(t *testing.T) {
	c := NewLRUCache(2)
	c.Put("a", 1)
	c.Put("b", 2)
	c.Get("a")        // "a" is now most recently used; "b" is LRU
	c.Put("c", 3)      // should evict "b"

	if _, ok := c.Get("b"); ok {
		t.Fatal("expected b to be evicted")
	}
	if _, ok := c.Get("a"); !ok {
		t.Fatal("expected a to still be present")
	}
	if _, ok := c.Get("c"); !ok {
		t.Fatal("expected c to be present")
	}
}

func TestLRUCache_PutExistingKeyUpdatesValueAndRecency(t *testing.T) {
	c := NewLRUCache(2)
	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("a", 100) // update + refresh recency
	c.Put("c", 3)   // should evict "b", not "a"

	if v, ok := c.Get("a"); !ok || v != 100 {
		t.Fatalf("expected a=100, got %v, ok=%v", v, ok)
	}
	if _, ok := c.Get("b"); ok {
		t.Fatal("expected b to have been evicted")
	}
}

func TestLRUCache_ConcurrentAccessDoesNotRace(t *testing.T) {
	c := NewLRUCache(16)
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := string(rune('a' + i%16))
			c.Put(key, i)
			c.Get(key)
		}(i)
	}
	wg.Wait()

	if c.Len() > 16 {
		t.Fatalf("cache exceeded capacity: len=%d", c.Len())
	}
}
