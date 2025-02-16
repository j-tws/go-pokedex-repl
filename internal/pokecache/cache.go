package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries map[string]cacheEntry
	Mux *sync.Mutex
}

type cacheEntry struct {
	CreatedAt time.Time
	Value []byte
}

func (c Cache) Add(key string, val []byte) {
	c.Entries[key] = cacheEntry{
		CreatedAt: time.Now(),
		Value: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheEntry, ok := c.Entries[key]
	
	if !ok {
		return nil, false
	}

	return cacheEntry.Value, true
}

func NewCache(interval time.Duration) Cache {
	c := Cache {
		Entries: map[string]cacheEntry{},
	}
	c.ReapLoop(interval)
	return c
}

func (c Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func(){
		for {
			<-ticker.C
			for key, cacheEntry := range c.Entries {
				timeDiff := time.Now().Sub(cacheEntry.CreatedAt)

				if timeDiff >= interval {
					delete(c.Entries, key)
				}
			}
		}
	}()
}