package main

import (
	"context"
	"file_manager/api"
	"file_manager/api/controllers"
	"file_manager/bootstrap"
	"file_manager/configs"
	"file_manager/internal/adapter/repositories"
	"file_manager/internal/common/log"
	"file_manager/internal/ports"
	"file_manager/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func init() {
	mode := "dev"
	configs.LoadConfigs(mode)
	//notice.InitNotification(conf.TelegramBotToken, conf.TelegramChatID)
}

func newGinEngine(logger log.Logging) (*gin.Engine, *gin.RouterGroup) {
	app := gin.New()
	app.Use(log.GinZap(logger.GetZap().Desugar()))
	return app, app.Group("")
}

func main() {
	ctx := context.Background()
	fx.New(
		fx.Supply(ctx),
		fx.Provide(log.NewLogger),
		fx.Invoke(log.RegisterGlobal),
		fx.Provide(bootstrap.InitDatabaseConnection),

		fx.Provide(repositories.NewUserRepository),
		fx.Provide(ports.InitUserRepositoryPort),

		fx.Provide(services.NewFileService),
		fx.Provide(services.NewAuthService),
		fx.Provide(services.NewRegisterService),

		fx.Provide(controllers.NewAppLiCationContext),

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
			})
		}),
	).Run()
}
