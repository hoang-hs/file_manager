package main

import (
	"context"
	"file_manager/src/api/bootstrap"
	"file_manager/src/common/log"
	"file_manager/src/common/notice"
	"file_manager/src/configs"
	"file_manager/src/monitor/jaeger"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	_ "net/http/pprof"
)

func init() {
	mode := "dev"
	configs.LoadConfigs(mode)
	cf := configs.Get()
	notice.InitNotification(cf.TelegramBotToken, cf.TelegramChatID)
	jaeger.InitJaeger()
}

func main() {
	fx.New(
		fx.Options(bootstrap.All()...),
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
