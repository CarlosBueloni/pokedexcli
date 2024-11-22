package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries map[string]cacheEntry
	sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) {

}

func (c *Cache) Add(key string, val []byte) {
	c.Lock()
	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Lock()
	entry, ok := c.Entries[key]
	c.Unlock()
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	tick := time.Tick(interval)
	for t := range tick {
		for _, entry := range c.Entries {
			if time.Since(entry.createdAt) > interval {

			}
		}
	}

}
