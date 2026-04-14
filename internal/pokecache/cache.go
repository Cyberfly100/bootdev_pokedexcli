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
	mu       sync.RWMutex
	cache    map[string]cacheEntry
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		interval: interval,
		cache:    make(map[string]cacheEntry),
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.cache[key]
	return value.val, exists
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	for key, entry := range c.cache {
		if now.Sub(entry.createdAt) > c.interval {
			delete(c.cache, key)
		}
	}
}
