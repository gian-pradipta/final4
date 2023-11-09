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

func (c *category) Create(newCategory dto.CreateCategoryRequest) error {
	var err error
	var entity entity.Category
	r := c.repo

	entity.Type = newCategory.Type
	err = r.Create(entity)
	return err
}
