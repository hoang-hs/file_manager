package api

import (
	"file_manager/api/controllers"
	"file_manager/api/middleware"
	"github.com/gin-gonic/gin"
)

type ControllerManager struct {
	FileController     *controllers.FileController
	LoginController    *controllers.LoginController
	RegisterController *controllers.RegisterController
}

func NewRouter(controllerManager *ControllerManager) *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/signup", controllerManager.RegisterController.SignUp)
	router.POST("/login", controllerManager.LoginController.Login)
	router.Use(middleware.RequiredJwtAuthentication())
	{
		router.GET("/tree/*path", controllerManager.FileController.Display)
		router.POST("/upload/*path", controllerManager.FileController.UploadFile)
		router.DELETE("/delete/*path", controllerManager.FileController.DeleteFile)
	}
	return router
}
