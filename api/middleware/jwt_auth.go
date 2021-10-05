package middleware

import "github.com/gin-gonic/gin"

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := bootstrap.TokenService.VerifyToken(c)
		if err != nil {
			c.JSON(err.GetHttpCode(), err.GetMessage())
			c.Abort()
			return
		}
		c.Next()
	}
}
