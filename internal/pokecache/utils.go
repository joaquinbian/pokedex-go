package pokecache

func deleteOldEntries(c *Cache) {
	for key, entry := range c.data {
		if entry.createdAt.After(entry.createdAt.Add(c.duration)) {
			delete(c.data, key)
		}
	}
}
