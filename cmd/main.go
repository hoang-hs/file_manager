package main

import (
	"context"
	"file_manager/api"
	"file_manager/api/controllers"
	"file_manager/api/database"
	"file_manager/api/services"
	"file_manager/configs"
	cacheAdapter "file_manager/internal/adapter/caching"
	"file_manager/internal/adapter/decorators"
	"file_manager/internal/adapter/repositories"
	"file_manager/internal/common/caching"
	"file_manager/internal/common/log"
	"file_manager/internal/common/notice"
	"file_manager/internal/ports"
	servicesImpl "file_manager/internal/services"
	"file_manager/jaeger"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func init() {
	mode := "dev"
	configs.LoadConfigs(mode)
	cf := configs.Get()
	notice.InitNotification(cf.TelegramBotToken, cf.TelegramChatID)
	jaeger.InitJaeger()
}

func newGinEngine(logger log.Logging) (*gin.Engine, *gin.RouterGroup) {
	app := gin.New()
	app.Use(log.GinZap(logger.GetZap().Desugar()))
	app.Use(log.RecoveryWithZap(logger.GetZap().Desugar()))
	return app, app.Group("")
}

func main() {
	ctx := context.Background()
	cf := configs.Get()
	fx.New(
		fx.Supply(ctx),
		fx.Supply(cf),
		fx.Provide(log.NewLogger),
		fx.Invoke(log.RegisterGlobal),

		fx.Supply(database.NewConnection(cf.DbDriver, cf.DbUser, cf.DbPassword,
			cf.DbPort, cf.DbHost, cf.DbName)),

		//bootstrap.LoadServices(),

		fx.Provide(repositories.NewBaseRepository),
		fx.Provide(repositories.NewUserCommandRepository),
		fx.Provide(repositories.NewUserQueryRepository),

		fx.Provide(cacheAdapter.NewInMemCache),
		fx.Provide(caching.InitCacheStrategy),

		fx.Provide(decorators.NewUserRepositoryDecorator),
		fx.Provide(ports.InitUserQueryRepositoryPort),
		fx.Provide(ports.InitUserCommandRepositoryPort),

		fx.Provide(servicesImpl.NewFileService),
		fx.Provide(servicesImpl.NewAuthService),
		fx.Provide(servicesImpl.NewRegisterService),

		fx.Provide(services.InitFileService),
		fx.Provide(services.InitAuthService),
		fx.Provide(services.InitRegisterService),

		fx.Provide(controllers.NewBaseController),
		fx.Provide(controllers.NewFileController),
		fx.Provide(controllers.NewRegisterController),
		fx.Provide(controllers.NewLoginController),

		fx.Provide(newGinEngine),

		fx.Invoke(api.InitRoutes),

		fx.Invoke(func(lc fx.Lifecycle, engine *gin.Engine) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go func() {
						if err := engine.Run(fmt.Sprintf("%s", configs.Get().ServerAddress)); err != nil {
							log.Fatalf("Cannot start application due by error [%v]", err)
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					log.Info("Stopping HTTP server.")
					return nil
					//return server.Shutdown(ctx)
				},
			})
		}),
	).Run()
}
