package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	rmu := &sync.RWMutex{}

	cache := Cache{
		mu:       rmu,
		duration: interval,
		data:     make(map[string]cacheEntry),
	}
	go cache.reapLoop()

	return cache

}

func (c *Cache) Add(key string, val []byte) {

	c.mu.RLock()

	c.data[key] = cacheEntry{
		createdAt: time.Time{},
		val:       val,
	}

	defer c.mu.RUnlock()
}

func (c Cache) Get(key string) ([]byte, bool) {

	c.mu.RLock()

	defer c.mu.RUnlock()

	data, ok := c.data[key]

	if !ok {
		return []byte{}, false
	}

	return data.val, true

}

func (c *Cache) reapLoop() {
	t := time.NewTicker(c.duration)
	for {
		<-t.C
		deleteOldEntries(c)
	}

}
