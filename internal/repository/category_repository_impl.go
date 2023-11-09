package repository

import (
	"database/sql"
	"final2/internal/entity"
	"time"
)

type category struct {
	db *sql.DB
}

func NewCategory(db *sql.DB) Category {
	var c category
	c.db = db
	return &c
}

func (c *category) Create(newCategory entity.Category) error {
	var err error
	db := c.db

	_, err = db.Exec("INSERT INTO category (type, created_at, updated_at) VALUES (?, ?, ?)", newCategory.Type, time.Now(), time.Now())
	return err
}
