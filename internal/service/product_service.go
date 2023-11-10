package service

import "final2/internal/dto"

type Product interface {
	GetByCategory(category int) ([]dto.GetProductsResponse, error)
}
