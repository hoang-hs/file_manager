package caching

import (
	"file_manager/internal/adapter/caching"
	"time"
)

type CacheStrategy interface {
	Get(key string) (interface{}, bool)
	Set(key string, data interface{}, ttl time.Duration)
}

func InitCacheStrategy(cacheImpl *caching.InMemCache) CacheStrategy {
	return cacheImpl
}
