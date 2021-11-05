package bootstrap

import (
	"database/sql"
	"file_manager/configs"
	log2 "file_manager/internal/common/log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func InitDatabaseConnection() *sql.DB {
	cf := configs.Get()
	dbDriver := cf.DbDriver
	dbUser := cf.DbUser
	dbPassword := cf.DbPassword
	dbPort := cf.DbPort
	dbHost := cf.DbHost
	dbName := cf.DbName
	return newConnection(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName)
}

func newConnection(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName string) *sql.DB {
	DbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	conn, err := sql.Open(dbDriver, DbUrl)
	if err != nil {
		log2.Fatalf("[Can not connect to database %s]: %s\n", dbDriver, err)
	} else {
		log2.Infof("Connected to database: %s\n", dbDriver)
	}
	return conn
}
