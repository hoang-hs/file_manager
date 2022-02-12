package bootstrap

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func NewConnection(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName string) *sql.DB {
	DbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	conn, err := sql.Open(dbDriver, DbUrl)
	if err != nil {
		log.Fatalf("Can not connect to driver: [%s], database: [%s], err: [%s]\n", dbDriver, dbName, err)
	} else {
		log.Printf("Connected to driver: %s, database: %s\n", dbDriver, dbName)
	}
	return conn
}
