package bootstrap

import (
	"file_manager/configs"
	cacheAdapter "file_manager/internal/adapter/caching"
	"file_manager/internal/adapter/decorators"
	"file_manager/internal/adapter/repositories"
	"file_manager/internal/common/caching"
	"file_manager/internal/ports"
	"file_manager/internal/services"
	"go.uber.org/fx"
)

func LoadServices() fx.Option {
	cf := configs.Get()
	env := cf.AppEnv
	return fx.Options(
		fx.Supply(newConnection(cf.DbDriver, cf.DbUser, cf.DbPassword,
			cf.DbPort, cf.DbHost, cf.DbName)),
		fx.Provide(repositories.NewBaseRepository),
		fx.Provide(repositories.NewUserCommandRepository),
		fx.Provide(repositories.NewUserQueryRepository),

		fx.Provide(cacheAdapter.NewInMemCache),
		fx.Provide(caching.InitCacheStrategy),

		fx.Supply(env),

		fx.Provide(decorators.NewUserRepositoryDecorator),
		fx.Provide(ports.InitUserQueryRepositoryPort),
		fx.Provide(ports.InitUserCommandRepositoryPort),

		fx.Provide(services.NewFileService),
		fx.Provide(services.NewAuthService),
		fx.Provide(services.NewRegisterService),
	)

}
