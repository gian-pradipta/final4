package repository

import "final2/internal/entity"

type Transaction interface {
	Create(transaction entity.Transaction) error
	GetUserProduct(userEmail string, productId int) (entity.User, entity.Product, error)
}
