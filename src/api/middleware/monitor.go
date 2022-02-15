package middleware

import (
	"file_manager/src/common/pubsub"
	"file_manager/src/core/events"
	"github.com/gin-gonic/gin"
	"time"
)

func SendRequestEvent() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		now := time.Now()
		c.Next()
		took := time.Since(now).Milliseconds()
		pubsub.Publish(events.NewRequestEvent(path, took))
	}
}
