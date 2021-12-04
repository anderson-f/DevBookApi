package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Connect open connection to database and returns
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringDatabaseConnection)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
