package repository

import "final2/internal/entity"

type Product interface {
	GetByCategory(category int) ([]entity.Product, error)
	Create(newProduct entity.Product) (int, error)
	Get(id int) (entity.Product, error)
	GetAll() ([]entity.Product, error)
}
