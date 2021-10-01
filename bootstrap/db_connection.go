package bootstrap

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func InitDatabaseConnection() *sql.DB {
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	return newConnection(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName)
}

func newConnection(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName string) *sql.DB {
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	conn, err := sql.Open(dbDriver, DBURL)
	if err != nil {
		log.Fatalf("[Can not connect to database %s]: %s\n", dbDriver, err)
	} else {
		log.Printf("[Connected to database] %s\n", dbDriver)
	}

	return conn
}
