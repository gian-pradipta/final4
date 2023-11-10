package repository

import "final2/internal/entity"

type Product interface {
	GetByCategory(category string) ([]entity.Product, error)
}
