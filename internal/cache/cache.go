package cache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt	time.Time
	val				[]byte
}

type Cache struct {
	data	map[string]CacheEntry
	mu		sync.Mutex
}

func New(interval time.Duration) *Cache {
	cache := &Cache{
		data: make(map[string]CacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.data[key] = CacheEntry{
		createdAt:	time.Now(),
		val:				val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if entry, ok := cache.data[key]; ok {
		return entry.val, true
	} else {
		return nil, false
	}
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for t := range ticker.C {
		for k, v := range cache.data {
			cache.mu.Lock()
			if t.Sub(v.createdAt) > interval {
				delete(cache.data, k)
			}
			cache.mu.Unlock()
		}
	}
}
