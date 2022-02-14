package bootstrap

import (
	"file_manager/src/api/controllers"
	"file_manager/src/api/registry"
	"go.uber.org/fx"
)

func LoadControllers() []fx.Option {
	return []fx.Option{
		fx.Provide(registry.NewFileService),
		fx.Provide(registry.NewRegisterService),
		fx.Provide(registry.NewAuthService),

		fx.Provide(controllers.NewBaseController),
		fx.Provide(controllers.NewFileController),
		fx.Provide(controllers.NewRegisterController),
		fx.Provide(controllers.NewLoginController),
	}
}
