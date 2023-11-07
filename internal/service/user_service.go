package service

import (
	"final2/internal/dto"
)

type User interface {
	Create(newUser dto.CreateUserRequest) error
	Login(newUser dto.LoginUserRequest) error
	TopUp(user dto.UpdateBalanceRequest, email string, group string) error
}
