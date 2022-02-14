package bootstrap

import (
	"file_manager/src/api/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func LoadEngine() []fx.Option {
	return []fx.Option{
		fx.Provide(gin.New),
		fx.Invoke(router.RegisterHandler),
		fx.Invoke(router.RegisterGinRouters),
	}
}
