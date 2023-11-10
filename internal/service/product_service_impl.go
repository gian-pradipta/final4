package service

import (
	"final2/internal/dto"
	"final2/internal/entity"
	"final2/internal/repository"
)

type product struct {
	repo repository.Product
}

func toGetProductResponse(entity entity.Product) dto.GetProductsResponse {
	var product dto.GetProductsResponse
	product.Id = entity.Id
	product.Title = entity.Title
	product.Price = entity.Price
	product.Stock = entity.Stock
	product.CreatedAt = entity.CreatedAt
	product.UpdatedAt = entity.UpdatedAt
	return product
}

func NewProduct(repo repository.Product) Product {
	var p product
	p.repo = repo
	return &p
}

func (p *product) GetByCategory(category int) ([]dto.GetProductsResponse, error) {
	var products []dto.GetProductsResponse = make([]dto.GetProductsResponse, 0)
	var err error
	r := p.repo

	entities, err := r.GetByCategory(category)
	for _, entity := range entities {
		product := toGetProductResponse(entity)
		products = append(products, product)
	}
	return products, err
}
