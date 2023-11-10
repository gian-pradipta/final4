package repository

import (
	"database/sql"
	"final2/internal/entity"

	"github.com/go-playground/validator/v10"
)

type product struct {
	db *sql.DB
	v  *validator.Validate
}

func NewProduct(db *sql.DB, v *validator.Validate) Product {
	var repo product
	repo.db = db
	repo.v = v
	return &repo
}

func (p *product) GetByCategory(category string) ([]entity.Product, error) {
	var products []entity.Product = make([]entity.Product, 1)
	var err error
	db := p.db

	rows, err := db.Query("SELECT * FROM product WHERE category = ?", category)
	if err != nil {
		return products, err
	}

	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.Id, &product.Title, &product.Price, &product.Stock, &product.CategoryId, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}
	return products, err
}
