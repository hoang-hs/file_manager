package bootstrap

import (
	"database/sql"
	"file_manager/src/adapter/database/repositories"
	"file_manager/src/configs"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/fx"
	"log"
)

func LoadRepositories(cf *configs.Config) []fx.Option {
	return []fx.Option{
		fx.Supply(newConnection(cf.DbDriver, cf.DbUser, cf.DbPassword,
			cf.DbPort, cf.DbHost, cf.DbName)),

		fx.Provide(repositories.NewBaseRepository),
		fx.Provide(repositories.NewUserCommandRepository),
		fx.Provide(repositories.NewUserQueryRepository),
	}
}

func newConnection(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName string) *sql.DB {
	DbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	conn, err := sql.Open(dbDriver, DbUrl)
	if err != nil {
		log.Fatalf("Can not connect to mysql, err: [%s], driver: [%s], database: [%s] \n", dbDriver, dbName, err)
	} else {
		log.Printf("Connected to driver: %s, database: %s\n", dbDriver, dbName)
	}
	return conn
}
