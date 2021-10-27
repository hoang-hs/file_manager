package main

import (
	"context"
	"file_manager/api"
	"file_manager/api/controllers"
	"file_manager/bootstrap"
	"file_manager/configs"
	"file_manager/internal/log"
	"file_manager/internal/repositories"
	"file_manager/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/fx"
)

func init() {
	mode := "dev"
	configs.LoadConfigs(mode)
}

func newGinEngine(logger log.Logging) (*gin.Engine, *gin.RouterGroup) {
	app := gin.New()
	//app.RedirectTrailingSlash = false
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
					log.Infof("Application will be served at %d. Service path: %s",
						8080, "/api")
					go func() {
						if err := engine.Run(fmt.Sprintf(":%d", 8080)); err != nil {
							log.Fatalf("Cannot start application due by error [%v]", err)
						}
					}()
					return nil
				},
			})
		}),
	).Run()
}
