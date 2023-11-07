package repository

import "final2/internal/entity"

type User interface {
	Create(newUser entity.User) error
	Login(newUser entity.User) error
	TopUp(user entity.User) error
}
