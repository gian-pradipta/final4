package service

import (
	"final2/internal/dto"
)

type User interface {
	Create(newUser dto.CreateUserRequest) (int, error)
	Login(newUser dto.LoginUserRequest) (string, error)
	TopUp(user dto.UpdateBalanceRequest, email string, group string) error
	GetCreateResponse(id int) (dto.CreateUserResponse, error)
}
