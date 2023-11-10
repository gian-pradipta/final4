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
