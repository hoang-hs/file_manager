package bootstrap

import (
	"file_manager/src/api/controllers"
	"go.uber.org/fx"
)

func LoadControllers() []fx.Option {
	return []fx.Option{
		fx.Provide(controllers.NewBaseController),
		fx.Provide(controllers.NewFileController),
		fx.Provide(controllers.NewRegisterController),
		fx.Provide(controllers.NewLoginController),
	}
}
