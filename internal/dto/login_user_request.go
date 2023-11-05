package dto

type LoginUserRequest struct {
	Email    string `json:"email" validator:"required,email"`
	Password string `json:"password" validator:"required,min=6"`
}
