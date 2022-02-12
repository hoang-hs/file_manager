package bootstrap

import (
	"context"
	"file_manager/src/adapter/caching"
	"file_manager/src/adapter/database/repositories"
	"file_manager/src/adapter/decorators"
	"file_manager/src/api"
	"file_manager/src/api/controllers"
	"file_manager/src/api/registry"
	caching2 "file_manager/src/common/caching"
	"file_manager/src/common/log"
	"file_manager/src/configs"
	"file_manager/src/core/ports"
	"file_manager/src/core/usecases"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func All() []fx.Option {
	ctx := context.Background()
	cf := configs.Get()
	return []fx.Option{
		fx.Supply(ctx),
		fx.Supply(cf),
		fx.Provide(log.NewLogger),
		fx.Invoke(log.RegisterGlobal),

		fx.Supply(NewConnection(cf.DbDriver, cf.DbUser, cf.DbPassword,
			cf.DbPort, cf.DbHost, cf.DbName)),

		fx.Provide(repositories.NewBaseRepository),
		fx.Provide(repositories.NewUserCommandRepository),
		fx.Provide(repositories.NewUserQueryRepository),

		fx.Provide(caching.NewInMemCache),
		fx.Provide(caching2.NewCacheStrategy),

		fx.Provide(decorators.NewUserRepositoryDecorator),
		fx.Provide(ports.NewUserQueryRepositoryPort),
		fx.Provide(ports.NewUserCommandRepositoryPort),

		fx.Provide(usecases.NewFileService),
		fx.Provide(usecases.NewAuthService),
		fx.Provide(usecases.NewRegisterService),

		fx.Provide(registry.NewFileService),
		fx.Provide(registry.NewRegisterService),
		fx.Provide(registry.NewAuthService),

		fx.Provide(controllers.NewBaseController),
		fx.Provide(controllers.NewFileController),
		fx.Provide(controllers.NewRegisterController),
		fx.Provide(controllers.NewLoginController),

		fx.Provide(gin.New),

		fx.Invoke(api.RegisterHandler),
		fx.Invoke(api.RegisterGinRouters),
	}
}
