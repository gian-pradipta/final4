package service

import (
	"final2/internal/dto"
	"final2/internal/helper/switchtype"
	"final2/internal/repository"
	"time"
)

type user struct {
	repo repository.User
}

func NewUser(userRepo repository.User) User {
	var s user
	s.repo = userRepo
	return &s
}

func (u *user) Create(newUser dto.CreateUserRequest) error {
	r := u.repo
	userEntity := switchtype.FromCreateUserRequestToUserEntity(newUser)
	userEntity.CreatedAt = time.Now()
	userEntity.UpdatedAt = time.Now()
	err := r.Create(userEntity)
	return err
}

func (u *user) Login(newUser dto.LoginUserRequest) error {
	var err error

	r := u.repo
	entity := switchtype.FromLoginUserRequestToEntityUser(newUser)
	err = r.Login(entity)
	return err
}

func (u *user) TopUp(user dto.UpdateBalanceRequest) error {
	var err error

	r := u.repo
	entity := switchtype.FromUpdateBalanceRequestToEntityUser(user)
	err = r.TopUp(entity)
	return err
}
