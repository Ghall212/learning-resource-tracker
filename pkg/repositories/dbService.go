package repositories

import "database/sql"

// DbService ...
type DbService interface {
	GetDB() (*sql.DB, error)
}
