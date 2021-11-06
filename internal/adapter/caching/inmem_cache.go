package caching

import (
	goCache "github.com/patrickmn/go-cache"
	"time"
)

const (
	TimeCache                    = 30 * time.Minute
	TimeCachePurgeExItemInMemory = 40 * time.Minute
)

type InMemCache struct {
	cache *goCache.Cache
}

func NewInMemCache() *InMemCache {
	return &InMemCache{
		cache: goCache.New(TimeCache, TimeCachePurgeExItemInMemory),
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
