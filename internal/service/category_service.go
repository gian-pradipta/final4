package service

import "final2/internal/dto"

type Category interface {
	Create(newCategory dto.CreateCategoryRequest) (int, error)
	Get(id int) (dto.CreateCategoryResponse, error)
}
