package caching

import (
	"file_manager/internal/adapter/caching"
	"time"
)

type CacheStrategy interface {
	Get(key string) (interface{}, bool)
	BatchGet(key []string) (interface{}, []string)
	Set(key string, data interface{}, ttl time.Duration)
}

func InitCacheStrategy(cache *caching.InMemCache) CacheStrategy {
	return cache
}
