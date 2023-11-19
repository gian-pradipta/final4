package service

import "final2/internal/dto"

type Transaction interface {
	Create(request dto.CreateTransactionRequest, userEmail string) (dto.CreateTransactionResponse, error)
}
