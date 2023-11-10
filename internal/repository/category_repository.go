package repository

import (
	"final2/internal/entity"
)

type Category interface {
	Create(newCategory entity.Category) (int, error)
	Get(id int) (entity.Category, error)
}
