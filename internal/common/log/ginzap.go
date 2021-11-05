package log

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GinZap(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		c.Next()
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("user_agent", c.Request.UserAgent()),
			//zap.String("time", time.Now().Format(timeFormat)),
		)
	}
}
