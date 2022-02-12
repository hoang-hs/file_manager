package caching

import (
	"file_manager/src/adapter/caching"
	"time"
)

type CacheStrategy interface {
	Get(key string) (interface{}, bool)
	Set(key string, data interface{}, ttl time.Duration)
}

func NewCacheStrategy(cacheImpl *caching.InMemCache) CacheStrategy {
	return cacheImpl
}
