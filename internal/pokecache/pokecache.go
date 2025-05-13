package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mutex   *sync.RWMutex
}

func (c Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, found := c.entries[key]; !found {
		c.entries[key] = cacheEntry{
			createdAt: time.Now(),
			val:       val,
		}
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	entry, found := c.entries[key]
	if found {
		return entry.val, true
	}
	return []byte{}, false
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func(ticker *time.Ticker) {
		for range ticker.C {
			for key, entry := range c.entries {
				if time.Now().After(entry.createdAt.Add(interval)) {
					delete(c.entries, key)
				}
			}
		}
	}(ticker)
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry, 1),
		mutex:   &sync.RWMutex{},
	}
	cache.reapLoop(interval)
	return cache
}
