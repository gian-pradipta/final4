package service

import "final2/internal/dto"

type Product interface {
	GetByCategory(category int) ([]dto.GetProductsResponse, error)
	Create(newProduct dto.CreateProductRequest) (int, error)
	Get(id int) (dto.GetProductsResponse, error)
	GetAll() ([]dto.GetProductsResponse, error)
}
