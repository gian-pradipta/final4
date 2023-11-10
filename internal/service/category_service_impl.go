package service

import (
	"final2/internal/dto"
	"final2/internal/entity"
	"final2/internal/repository"
)

type category struct {
	repo repository.Category
}

func NewCategory(repo repository.Category) Category {
	var ctg category
	ctg.repo = repo
	return &ctg
}

func toGetCategoryResponse(entity entity.CategoryWithProduct) dto.GetCategoriesResponse {
	var response dto.GetCategoriesResponse
	var products []dto.GetProductsResponse = make([]dto.GetProductsResponse, 0)
	response.Id = entity.Id
	response.Type = entity.Type
	response.SoldProductAmount = entity.SoldProductAmount
	response.CreatedAt = entity.CreatedAt
	response.UpdatedAt = entity.UpdatedAt
	for _, p := range entity.Products {
		product := toGetProductResponse(p)
		products = append(products, product)
	}
	response.Products = products
	return response
}

func (c *category) Create(newCategory dto.CreateCategoryRequest) (int, error) {
	var err error
	var entity entity.Category
	var id int
	r := c.repo

	entity.Type = newCategory.Type
	id, err = r.Create(entity)
	return id, err
}

func (c *category) Get(id int) (dto.CreateCategoryResponse, error) {
	var err error
	var entity entity.Category
	var category dto.CreateCategoryResponse
	r := c.repo
	entity, err = r.Get(id)
	category.Id = entity.Id
	category.Type = entity.Type
	category.SoldProductAmount = entity.SoldProductAmount
	category.CreatedAt = entity.CreatedAt
	return category, err
}

func (c *category) GetAll() ([]dto.GetCategoriesResponse, error) {
	var categories []dto.GetCategoriesResponse
	var err error
	repo := c.repo

	entities, err := repo.GetAll()
	if err != nil {
		return categories, err
	}
	for _, entity := range entities {
		category := toGetCategoryResponse(entity)
		categories = append(categories, category)
	}
	return categories, err
}

func (c *category) Update(newCategory dto.CreateCategoryRequest, id int) (int, error) {
	var err error
	var entity entity.Category
	r := c.repo

	entity.Type = newCategory.Type
	id, err = r.Update(entity, id)
	return id, err
}

func (c *category) Delete(id int) error {
	var err error
	repo := c.repo

	err = repo.Delete(id)
	return err
}
