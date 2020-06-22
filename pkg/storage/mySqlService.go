package storage

import (
	"database/sql"

	// Should be acceptable to not be in package main because this code always
	// registers a driver for this package, registering in main would
	// not separate concerns appropriately
	_ "github.com/go-sql-driver/mysql"
)

// MySQLService ...
type MySQLService struct {
	ConnectionString string
}

// GetDB returns an instance of the database
func (ms *MySQLService) GetDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", ms.ConnectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
