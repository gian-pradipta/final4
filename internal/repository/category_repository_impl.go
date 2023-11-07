package repository

import "database/sql"

type category struct {
	db *sql.DB
}

func NewCategory(db *sql.DB) Category {
	var c category
	c.db = db
	return db
}
