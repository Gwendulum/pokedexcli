package pokecache

import(
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt	time.Time
	value		[]byte
}
type Cache struct {
	cache 		map[string]cacheEntry
	mu 			sync.Mutex
	interval 	time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache: make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt:	time.Now(),
		value:		val,
	}	
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.cache[key]
	if !exists {
		return nil, false
	}
	return entry.value, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for now := range ticker.C {
		
		for key, entry := range c.cache {
			c.mu.Lock()
			if now.Sub(entry.createdAt) > c.interval {
				delete(c.cache, key)
			}	
			c.mu.Unlock()
		}
	}
}