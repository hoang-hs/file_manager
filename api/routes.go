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

type ControllerManager struct {
	FileController     *controllers.FileController
	LoginController    *controllers.LoginController
	RegisterController *controllers.RegisterController
}

func NewRouter(controllerManager *ControllerManager) *gin.Engine {
	router := gin.Default()
	router.POST("/signup", controllerManager.RegisterController.SignUp)
	router.POST("/login", controllerManager.LoginController.Login)
	router.GET("/tree/*path", controllers.MiddleWare(), controllerManager.FileController.Display)
	router.POST("/upload/*path", controllers.MiddleWare(), controllerManager.FileController.UploadFile)
	router.DELETE("/delete/*path", controllers.MiddleWare(), controllerManager.FileController.DeleteFile)
	return router
}
