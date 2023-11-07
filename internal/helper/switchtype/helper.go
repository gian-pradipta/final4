package switchtype

import (
	"final2/internal/dto"
	"final2/internal/entity"
)

func FromCreateUserRequestToUserEntity(dtoUser dto.CreateUserRequest) entity.User {
	var user entity.User
	user.Email = dtoUser.Email
	user.Fullname = dtoUser.Fullname
	user.Password = dtoUser.Password
	return user
}

func FromLoginUserRequestToEntityUser(dtoUser dto.LoginUserRequest) entity.User {
	var user entity.User
	user.Email = dtoUser.Email
	user.Fullname = "lorem"
	user.Password = dtoUser.Password
	user.Balance = 20
	return user
}

func FromUpdateBalanceRequestToEntityUser(dtoUser dto.UpdateBalanceRequest, email string, group string) entity.User {
	var user entity.User
	user.Balance = dtoUser.Balance
	user.Email = email
	user.Role = group
	return user
}
