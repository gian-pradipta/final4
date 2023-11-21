package service

import (
	"final2/internal/dto"
	"final2/internal/entity"
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

func (u *user) Create(newUser dto.CreateUserRequest) (int, error) {
	r := u.repo
	userEntity := switchtype.FromCreateUserRequestToUserEntity(newUser)
	userEntity.CreatedAt = time.Now()
	userEntity.UpdatedAt = time.Now()
	id, err := r.Create(userEntity)
	return id, err
}

func (u *user) Login(newUser dto.LoginUserRequest) (string, error) {
	var err error

	r := u.repo
	entity := switchtype.FromLoginUserRequestToEntityUser(newUser)
	group, err := r.Login(entity)
	return group, err
}

func (u *user) TopUp(user dto.UpdateBalanceRequest, email string, group string) error {
	var err error

	r := u.repo
	entity := switchtype.FromUpdateBalanceRequestToEntityUser(user, email, group)
	err = r.TopUp(entity)
	return err
}

func toGetMyTransactiobResponse(entity entity.CategoryWithProduct) dto.GetCategoriesResponse {
	var response dto.GetCategoriesResponse
	var products []dto.GetProductsResponse = make([]dto.GetProductsResponse, 0)
	response.Id = entity.Id
	response.Type = entity.Type
	response.SoldProductAmount = entity.SoldProductAmount
	response.CreatedAt = entity.CreatedAt
	response.UpdatedAt = entity.UpdatedAt
	for _, p := range entity.Products {
		product := toGetProductResponse(p)
		products = append(products, product)
	}
	response.Products = products
	return response
}

func (u *user) GetCreateResponse(id int) (dto.CreateUserResponse, error) {
	var err error
	var user dto.CreateUserResponse
	entity, err := u.repo.Get(id)
	if err != nil {
		return user, err
	}
	user.Balance = entity.Balance
	user.CreatedAt = entity.CreatedAt
	user.Fullname = entity.Fullname
	user.Id = entity.Id
	user.Email = entity.Email
	return user, err
}
