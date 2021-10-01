package api

import (
	"file_manager/api/controllers"
	"github.com/gin-gonic/gin"
)

/*
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
*/

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//router.Use(CORSMiddleware())

	router.POST("/signup", controllers.RegisterController.SignUp)
	router.POST("/login", controllers.Login)
	router.GET("/tree/*path", controllers.MiddleWare(), controllers.Display)
	router.POST("/upload/*path", controllers.MiddleWare(), controllers.UploadFile)
	router.DELETE("/delete/*path", controllers.MiddleWare(), controllers.DeleteFile)

	return router
}
