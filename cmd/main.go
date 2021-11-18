package main

import (
	"context"
	"file_manager/api"
	"file_manager/api/controllers"
	"file_manager/bootstrap"
	"file_manager/configs"
	"file_manager/internal/common/log"
	"file_manager/internal/common/notice"
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
	fx.New(
		fx.Supply(ctx),
		fx.Provide(log.NewLogger),
		fx.Invoke(log.RegisterGlobal),

		bootstrap.LoadServices(),

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
				OnStop: func(ctx context.Context) error {
					log.Info("Stopping HTTP server.")
					return nil
					//return server.Shutdown(ctx)
				},
			})
		}),
	).Run()
}
