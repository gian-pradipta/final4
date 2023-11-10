package repository

import "final2/internal/entity"

type Product interface {
	GetByCategory(category int) ([]entity.Product, error)
}
