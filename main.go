package main

import (
	"context"
	"file_manager/src/api/bootstrap"
	"file_manager/src/common/log"
	"file_manager/src/common/notice"
	"file_manager/src/configs"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	_ "net/http/pprof"
)

func init() {
	mode := "dev"
	configs.LoadConfigs(mode)
	cf := configs.Get()
	logger, err := log.NewLogger()
	if err != nil {
		panic(err)
	}
	log.RegisterGlobal(logger)
	notice.InitNotification(cf.TelegramBotToken, cf.TelegramChatID)
}

func main() {
	fx.New(
		fx.Provide(configs.Get),
		fx.Provide(log.GetGlobalLog),
		fx.Options(bootstrap.LoadRepositories()...),
		fx.Options(bootstrap.LoadUseCases()...),
		fx.Options(bootstrap.LoadControllers()...),
		fx.Options(bootstrap.LoadEngine()...),

		fx.Options(bootstrap.LoadGraphite()...),
		fx.Options(bootstrap.LoadListeners()...),
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
				},
			})
		}),
	).Run()
}
