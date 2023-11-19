package dto

type CreateUserRequest struct {
	Fullname string `json:"fullname" validate:"required,min=1"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
