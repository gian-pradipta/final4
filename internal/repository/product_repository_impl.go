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

func (p *product) Create(newProduct entity.Product) (int, error) {
	var err error
	var latestId int
	db := p.db

	result, err := db.Exec("INSERT INTO product (title, price, stock, category_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", newProduct.Title, newProduct.Price, newProduct.Stock, newProduct.CategoryId, time.Now(), time.Now())
	if err != nil {
		return latestId, err
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return latestId, err
	}
	latestId = int(id64)
	return latestId, err
}

func (p *product) Get(id int) (entity.Product, error) {
	var product entity.Product
	db := p.db
	var err error
	var updatedAt string
	var createdAt string
	rows := db.QueryRow("SELECT * FROM product WHERE id = ?", id)

	err = rows.Scan(&product.Id, &product.Title, &product.Price, &product.Stock, &product.CategoryId, &createdAt, &updatedAt)
	if err != nil {
		return product, err
	}
	product.CreatedAt, err = time.Parse("2006-01-02 15:04:05.9999999-07:00", createdAt)
	product.UpdatedAt, err = time.Parse("2006-01-02 15:04:05.9999999-07:00", updatedAt)
	return product, err

}

func (p *product) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	db := p.db
	var err error
	var updatedAt string
	var createdAt string
	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
		return products, err
	}
	defer rows.Close()
	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.Id, &product.Title, &product.Price, &product.Stock, &product.CategoryId, &createdAt, &updatedAt)
		product.CreatedAt, err = time.Parse("2006-01-02 15:04:05.9999999-07:00", createdAt)
		product.UpdatedAt, err = time.Parse("2006-01-02 15:04:05.9999999-07:00", updatedAt)
		products = append(products, product)
	}
	return products, err
}

func (p *product) Update(newProduct entity.Product, id int) (int, error) {
	var err error
	// fmt.Println(newProduct)
	_, err = p.Get(id)
	if err != nil {
		return 0, err
	}
	_, err = p.db.Exec("UPDATE product SET title = ?, price = ?, stock = ?, category_id = ?, updated_at = ? WHERE id = ?", newProduct.Title, newProduct.Price, newProduct.Stock, newProduct.CategoryId, time.Now(), id)
	return id, err
}

func (p *product) Delete(id int) error {
	var err error
	_, err = p.Get(id)
	if err != nil {
		return err
	}
	_, err = p.db.Exec("DELETE FROM product WHERE id = ?", id)
	return err
}
