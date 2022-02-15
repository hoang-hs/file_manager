package bootstrap

import (
	"database/sql"
	"file_manager/src/adapter/database/repositories"
	"file_manager/src/common/log"
	"file_manager/src/configs"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/fx"
)

func LoadRepositories() []fx.Option {
	return []fx.Option{
		fx.Provide(newConnection),
		fx.Provide(repositories.NewBaseRepository),
		fx.Provide(repositories.NewUserCommandRepository),
		fx.Provide(repositories.NewUserQueryRepository),
	}
}

func newConnection(cf *configs.Config) *sql.DB {
	DbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cf.DbUser, cf.DbPassword,
		cf.DbPort, cf.DbHost, cf.DbName)
	conn, err := sql.Open(cf.DbDriver, DbUrl)
	if err != nil {
		log.Fatalf("Can not connect to mysql, err: [%s], driver: [%s], database: [%s] \n", cf.DbDriver, cf.DbName, err)
	} else {
		log.Infof("Connected to driver: %s, database: %s\n", cf.DbDriver, cf.DbName)
	}
	return conn
}
