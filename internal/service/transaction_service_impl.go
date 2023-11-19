package service

import (
	"errors"
	"final2/internal/dto"
	"final2/internal/entity"
	"final2/internal/repository"
	"time"
)

type transaction struct {
	repo repository.Transaction
}

func NewTransaction(repo repository.Transaction) Transaction {
	var t transaction
	t.repo = repo
	return &t
}

func validateTransaction(transaction entity.Transaction, product entity.Product, user entity.User) bool {
	var valid bool = true
	if product.Stock < transaction.Quantity {
		return false
	}
	if product.Price*transaction.Quantity > user.Balance {
		return false
	}
	return valid
}

func (t *transaction) Create(request dto.CreateTransactionRequest, userEmail string) (dto.CreateTransactionResponse, error) {

	repo := t.repo
	var err error
	var response dto.CreateTransactionResponse
	var transaction entity.Transaction
	user, product, err := repo.GetUserProduct(userEmail, request.ProductId)

	transaction.ProductId = request.ProductId
	transaction.UserId = user.Id
	transaction.Quantity = request.Quantity
	transaction.TotalPrice = request.Quantity * product.Price
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()
	// fmt.Println(transaction)
	valid := validateTransaction(transaction, product, user)
	if !valid {
		err = errors.New("Invalid Transaction")
		return response, err
	}

	err = t.repo.Create(transaction)
	response.Message = "You have successfully purchased the product"
	response.TransactionBill.ProductTitle = product.Title
	response.TransactionBill.Quantity = transaction.Quantity
	response.TransactionBill.TotalPrice = transaction.TotalPrice

	return response, err
}

func toGetMyTransactionsResponse(entities []entity.TransactionWithProduct) []dto.GetTransactionResponse {
	// var response dto.GetTransactionResponse
	var myTransactions []dto.GetTransactionResponse

	for _, entity := range entities {
		var t dto.GetTransactionResponse
		t.Id = entity.Id
		t.ProductId = entity.ProductId
		t.UserId = entity.UserId
		t.Quantity = entity.Quantity
		t.TotalPrice = entity.TotalPrice
		t.Product.Id = entity.Product.Id
		t.Product.Title = entity.Product.Title
		t.Product.Price = entity.Product.Price
		t.Product.Stock = entity.Product.Stock
		t.Product.CategoryId = entity.Product.CategoryId
		t.Product.CreatedAt = entity.Product.CreatedAt
		t.Product.UpdatedAt = entity.Product.UpdatedAt
		myTransactions = append(myTransactions, t)
	}
	return myTransactions
}

func toGetAllTransactionsResponse(entities []entity.TransactionWithProductUser) []dto.GetAllTransactionResponse {
	var myTransactions []dto.GetAllTransactionResponse

	for _, entity := range entities {
		var t dto.GetAllTransactionResponse
		t.Id = entity.Id
		t.ProductId = entity.ProductId
		t.UserId = entity.UserId
		t.Quantity = entity.Quantity
		t.TotalPrice = entity.TotalPrice
		t.Product.Id = entity.Product.Id
		t.Product.Title = entity.Product.Title
		t.Product.Price = entity.Product.Price
		t.Product.Stock = entity.Product.Stock
		t.Product.CategoryId = entity.Product.CategoryId
		t.Product.CreatedAt = entity.Product.CreatedAt
		t.Product.UpdatedAt = entity.Product.UpdatedAt
		t.User.Id = entity.User.Id
		t.User.Email = entity.User.Email
		t.User.Fullname = entity.User.Fullname
		t.User.Balance = entity.User.Balance
		t.User.CreatedAt = entity.User.CreatedAt
		t.User.UpdatedAt = entity.User.UpdatedAt
		myTransactions = append(myTransactions, t)
	}
	return myTransactions
}

func (t *transaction) GetMyTransactions(email string) ([]dto.GetTransactionResponse, error) {
	var myTransactions []dto.GetTransactionResponse
	var err error
	repo := t.repo
	entities, err := repo.GetMyTransactions(email)
	if err != nil {
		return myTransactions, err
	}
	myTransactions = toGetMyTransactionsResponse(entities)

	return myTransactions, err
}

func (t *transaction) GetAllTransactions() ([]dto.GetAllTransactionResponse, error) {
	var transactions []dto.GetAllTransactionResponse
	var err error
	repo := t.repo

	entities, err := repo.GetAllTransactions()
	if err != nil {
		return transactions, err
	}
	transactions = toGetAllTransactionsResponse(entities)
	return transactions, err
}
