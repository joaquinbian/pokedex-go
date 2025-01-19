package pokecache

import (
	"time"
)

func deleteOldEntries(c *Cache) {
	for key, entry := range c.data {
		if time.Since(entry.createdAt) > c.duration {
			delete(c.data, key)
		}
	}
}
