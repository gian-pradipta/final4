package repository

import "final2/internal/entity"

type Transaction interface {
	Create(transaction entity.Transaction) error
	GetUserProduct(userEmail string, productId int) (entity.User, entity.Product, error)
	GetMyTransactions(userEmail string) ([]entity.TransactionWithProduct, error)
	GetAllTransactions() ([]entity.TransactionWithProductUser, error)
}
