package service

import (
	"final2/internal/dto"
	"final2/internal/entity"
	"final2/internal/repository"
)

type product struct {
	repo repository.Product
}

func toEntity(product dto.CreateProductRequest) entity.Product {
	var entity entity.Product
	entity.CategoryId = product.CategoryId
	entity.Title = product.Title
	entity.Price = product.Price
	entity.Stock = product.Stock
	return entity

}

func toGetProductResponse(entity entity.Product) dto.GetProductsResponse {
	var product dto.GetProductsResponse
	product.Id = entity.Id
	product.Title = entity.Title
	product.Price = entity.Price
	product.Stock = entity.Stock
	product.CategoryId = entity.CategoryId
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

func (p *product) Create(newProduct dto.CreateProductRequest) (int, error) {
	var err error
	entity := toEntity(newProduct)
	id, err := p.repo.Create(entity)
	return id, err
}

func (p *product) Get(id int) (dto.GetProductsResponse, error) {
	var err error
	var entity entity.Product
	var product dto.GetProductsResponse
	r := p.repo
	entity, err = r.Get(id)
	product = toGetProductResponse(entity)

	return product, err
}

func (p *product) GetAll() ([]dto.GetProductsResponse, error) {
	var err error
	var products []dto.GetProductsResponse
	var entities []entity.Product
	r := p.repo
	entities, err = r.GetAll()
	if err != nil {
		return products, err
	}
	for _, entity := range entities {
		product := toGetProductResponse(entity)
		products = append(products, product)
	}

	return products, err
}

func (p *product) Update(id int, product dto.CreateProductRequest) (int, error) {
	var err error

	entity := toEntity(product)
	_, err = p.repo.Update(entity, id)
	return id, err
}

func (p *product) Delete(id int) error {
	var err error
	err = p.repo.Delete(id)
	return err

}
