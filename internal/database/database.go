package database

import (
	"database/sql"

	_ "github.com/glebarez/sqlite"
)

const DRIVER_NAME string = "sqlite"
const DATABASE string = "todos.db"

func New() (*sql.DB, error) {
	var err error
	var db *sql.DB
	db, err = sql.Open(DRIVER_NAME, DATABASE)
	db.Exec("PRAGMA foreign_keys = ON")
	return db, err
}
