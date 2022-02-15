package bootstrap

import (
	"file_manager/src/adapter/caching"
	"file_manager/src/adapter/decorators"
	caching2 "file_manager/src/common/caching"
	"file_manager/src/core/usecases"
	"go.uber.org/fx"
)

func LoadUseCases() []fx.Option {
	return []fx.Option{
		fx.Provide(caching.NewInMemCache),
		fx.Provide(caching2.NewCacheStrategy),

		fx.Provide(decorators.NewUserRepositoryDecorator),

		fx.Provide(usecases.NewFileUseCase),
		fx.Provide(usecases.NewAuthUseCase),
		fx.Provide(usecases.NewRegisterService),
		fx.Provide(usecases.NewUpdateMetricUseCase),
	}
}
