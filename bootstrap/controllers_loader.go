package bootstrap

import (
	"file_manager/api"
	"file_manager/api/controllers"
)

func LoadControllers(appContext *controllers.ApplicationContext) *api.ControllerManager {
	fileController := controllers.NewFileController(appContext)
	loginController := controllers.NewLoginController(appContext)
	registerController := controllers.NewRegisterController(appContext)
	return &api.ControllerManager{
		FileController:     fileController,
		LoginController:    loginController,
		RegisterController: registerController,
	}
}
