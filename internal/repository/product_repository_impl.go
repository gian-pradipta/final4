package repository

import (
	"database/sql"
	"final2/internal/entity"
	"time"
)

type product struct {
	db *sql.DB
}

func NewProduct(db *sql.DB) Product {
	var repo product
	repo.db = db
	return &repo
}

func (p *product) GetByCategory(category int) ([]entity.Product, error) {
	var products []entity.Product = make([]entity.Product, 0)
	var err error
	var updatedAt string
	var createdAt string
	db := p.db

	rows, err := db.Query("SELECT * FROM product WHERE category_id = ?", category)
	if err != nil {
		return products, err
	}

	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.Id, &product.Title, &product.Price, &product.Stock, &product.CategoryId, &createdAt, &updatedAt)
		product.CreatedAt, err = time.Parse("2006-01-02 15:04:05.9999999-07:00", createdAt)
		product.UpdatedAt, err = time.Parse("2006-01-02 15:04:05.9999999-07:00", updatedAt)
		if err != nil {
			return products, err
		}
		if product.Id != 0 {

			products = append(products, product)
		}
	}
	return products, err
}
