package api

import (
	"file_manager/api/controllers"
	"file_manager/api/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type RoutersIn struct {
	fx.In
	FileController     *controllers.FileController
	LoginController    *controllers.LoginController
	RegisterController *controllers.RegisterController
}

func InitRoutes(group *gin.RouterGroup, in RoutersIn) {
	registerPublicRoutes(group, in)
	group.Use(middleware.RequiredJwtAuthentication())
	{
		registerProtectedRoutes(group, in)
	}
}

func registerPublicRoutes(r *gin.RouterGroup, in RoutersIn) {
	r.POST("/signup", in.RegisterController.SignUp)
	r.POST("/login", in.LoginController.Login)
}

func registerProtectedRoutes(r *gin.RouterGroup, in RoutersIn) {
	r.GET("/tree/*path", in.FileController.Display)
	r.POST("/upload/*path", in.FileController.UploadFile)
	r.DELETE("/delete/*path", in.FileController.DeleteFile)
}
