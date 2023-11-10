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

func (c *category) Create(newCategory entity.Category) (int, error) {
	var err error
	var id int
	db := c.db

	result, err := db.Exec("INSERT INTO category (type, created_at, updated_at) VALUES (?, ?, ?)", newCategory.Type, time.Now(), time.Now())
	if err != nil {
		return id, err
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return id, err
	}

	id = int(id64)
	return id, err
}

func (c *category) Get(id int) (entity.Category, error) {
	db := c.db
	var err error
	var category entity.Category
	var updatedAt string
	var createdAt string
	rows := db.QueryRow("SELECT * FROM category WHERE id = ?", id)

	err = rows.Scan(&category.Id, &category.Type, &category.SoldProductAmount, &createdAt, &updatedAt)
	category.CreatedAt, err = time.Parse("2006-01-02 15:04:05.9999999-07:00", createdAt)
	category.UpdatedAt, err = time.Parse("2006-01-02 15:04:05.9999999-07:00", updatedAt)
	return category, err

}
