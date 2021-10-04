package main

import (
	"file_manager/api"
	"file_manager/bootstrap"
	"file_manager/configs"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs.LoadConfigs()
	dbConnection := bootstrap.InitDatabaseConnection()
	defer func() {
		_ = dbConnection.Close()
	}()
	appContext := bootstrap.LoadServices(dbConnection)
	controller := bootstrap.LoadControllers(appContext)
	server := api.NewServer(api.NewRouter(controller))
	server.Run(configs.Get().Port)
}
