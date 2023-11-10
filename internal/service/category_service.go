package service

import "final2/internal/dto"

type Category interface {
	Create(newCategory dto.CreateCategoryRequest) (int, error)
	Get(id int) (dto.CreateCategoryResponse, error)
	GetAll() ([]dto.GetCategoriesResponse, error)
	Update(newCategory dto.CreateCategoryRequest, id int) (int, error)
}
