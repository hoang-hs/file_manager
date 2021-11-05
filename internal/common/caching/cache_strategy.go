package caching

import "time"

type CacheStrategy interface {
	Get(key string) (interface{}, bool)
	BatchGet(key []string) (interface{}, []string)
	Set(key string, data interface{}, ttl time.Duration)
}
