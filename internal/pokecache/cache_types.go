package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data     map[string]cacheEntry
	duration time.Duration
	mu       *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
