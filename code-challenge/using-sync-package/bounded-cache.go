package using_sync_package

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type entry struct {
	value any
	time  time.Time
}

type Cache struct {
	size int
	ttl  time.Duration

	mu      sync.RWMutex
	entries map[string]entry
	cancel  context.CancelFunc
}

func New(size int, ttl time.Duration) (*Cache, error) {
	if size <= 0 {
		return nil, fmt.Errorf("size must be greater than zero - %d", size)
	}

	if ttl <= 0 {
		return nil, fmt.Errorf("ttl must be greater than zero - %v", ttl)
	}

	ctx, cancel := context.WithCancel(context.Background())
	c := &Cache{
		size:    size,
		ttl:     ttl,
		entries: make(map[string]entry),
		cancel:  cancel,
	}

	go c.janitor(ctx)
	return c, nil
}

func (c *Cache) Close() {
	c.cancel()
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return v.value, true
}

func (c *Cache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == len(c.entries) {
		c.popOne()
	}

	ve := entry{value, time.Now()}
	c.entries[key] = ve
}

func (c *Cache) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	nKeys := make([]string, 0, len(c.entries))
	for k := range c.entries {
		nKeys = append(nKeys, k)
	}

	return nKeys
}

func (c *Cache) popOne() {
	tmK, tmT := "", time.Now()

	for k, v := range c.entries {
		if v.time.Before(tmT) {
			tmK, tmT = k, v.time
		}
	}

	delete(c.entries, tmK)
}

func (c *Cache) janitor(ctx context.Context) {
	ticker := time.NewTicker(c.ttl)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.clueanUp()
		case <-ctx.Done():
			return
		}
	}
}

func (c *Cache) clueanUp() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for k, v := range c.entries {
		if now.Sub(v.time) > c.ttl {
			delete(c.entries, k)
		}
	}
}
