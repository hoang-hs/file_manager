package caching

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type InMemCache struct {
	cache *cache.Cache
}

func NewInMemCache(cache *cache.Cache) *InMemCache {
	return &InMemCache{
		cache: cache,
	}
}

func (i *InMemCache) Get(key string) (interface{}, bool) {
	return i.cache.Get(key)
}

func (i *InMemCache) Set(key string, data interface{}, ttl time.Duration) {
	i.cache.Set(key, data, ttl)
}

// TODO implement BatchGet
func (i *InMemCache) BatchGet(key []string) (interface{}, []string) {
	panic("implement me")
}
