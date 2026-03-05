
package creative

import "sync"

type Cache struct {

	mu sync.RWMutex
	store map[string]string
}

func NewCache() *Cache {

	return &Cache{
		store: make(map[string]string),
	}
}

func (c *Cache) Get(id string) (string,bool) {

	c.mu.RLock()
	defer c.mu.RUnlock()

	v,ok := c.store[id]

	return v,ok
}

func (c *Cache) Set(id string, vast string) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.store[id] = vast
}
