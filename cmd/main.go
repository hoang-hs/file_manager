package main

import (
	"file_manager/api"
	"file_manager/bootstrap"
	"file_manager/configs"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs.LoadEnv()
	configs.LoadConfigs()

	dbConnection := bootstrap.InitDatabaseConnection()
	defer func() {
		_ = dbConnection.Close()
	}()
	bootstrap.LoadServices(dbConnection)
	server := api.NewServer(api.NewRouter())
	server.Run(":8080")
}
