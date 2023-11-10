package repository

import (
	"final2/internal/entity"
)

type Category interface {
	Create(newCategory entity.Category) (int, error)
	Get(id int) (entity.Category, error)
	GetAll() ([]entity.CategoryWithProduct, error)
	Update(newCategory entity.Category, id int) (int, error)
}
