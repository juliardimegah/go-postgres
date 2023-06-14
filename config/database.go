package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DBConnection() (*sql.DB, error) {
	dbDriver := "postgres"
	dbUser := "postgres"
	dbPass := "try"
	dbName := "member"
	dbHost := "localhost"
	dbPort := "5000"

	connStr := "postgres://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName
	db, err := sql.Open(dbDriver, connStr)
	return db, err
}
