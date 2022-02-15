package router

import (
	"file_manager/src/api/controllers"
	"file_manager/src/api/middleware"
	"file_manager/src/common/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type RoutersIn struct {
	fx.In
	Engine             *gin.Engine
	FileController     *controllers.FileController
	LoginController    *controllers.LoginController
	RegisterController *controllers.RegisterController
}

func RegisterHandler(engine *gin.Engine, logger log.Logging) {
	engine.Use(middleware.SendRequestEvent())
	engine.Use(log.GinZap(logger.GetZap().Desugar()))
	engine.Use(log.RecoveryWithZap(logger.GetZap().Desugar()))
}

func RegisterGinRouters(in RoutersIn) {
	group := in.Engine.Group("/")
	registerPublicRouters(group, in)
	group.Use(middleware.RequiredJwtAuthentication())
	{
		registerProtectedRouters(group, in)
	}
}

func registerPublicRouters(r *gin.RouterGroup, in RoutersIn) {
	r.POST("/signup", in.RegisterController.SignUp)
	r.POST("/login", in.LoginController.Login)
}

func registerProtectedRouters(r *gin.RouterGroup, in RoutersIn) {
	r.GET("/tree", in.FileController.Display)
	r.POST("/upload", in.FileController.UploadFile)
	r.DELETE("/delete", in.FileController.DeleteFile)
}
