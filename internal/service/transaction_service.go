package service

import "final2/internal/dto"

type Transaction interface {
	GetMyTransactions(email string) ([]dto.GetTransactionResponse, error)
	Create(request dto.CreateTransactionRequest, userEmail string) (dto.CreateTransactionResponse, error)
	GetAllTransactions() ([]dto.GetAllTransactionResponse, error)
}
