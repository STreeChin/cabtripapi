package model

import (
	"github.com/allegro/bigcache"
	"log"
	"sync"
	"time"
)

type Cache struct {
	bc *bigcache.BigCache
}

var cache *Cache
var once sync.Once

func GetCacheInstance() *Cache {
	once.Do(func() {
		cache = new(Cache)
		cache.bc, _ = bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	})
	return cache
}

func (c *Cache) GetCache(key string) (string, error) {
	res, err := c.bc.Get(key)
	return string(res), err
}

func (c *Cache) SetCache(key string, count string) {
	err := c.bc.Set(key, []byte(count))
	if err != nil {
		log.Fatal(err)
	}
}
func (c *Cache) ResetCache() {
	err := c.bc.Reset()
	if err != nil {
		log.Fatal(err)
	}
}
