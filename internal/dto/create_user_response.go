package dto

import "time"

type CreateUserResponse struct {
	Id        int       `json:"id"`
	Fullname  string    `json:"fullname,min=1"`
	Email     string    `json:"email"`
	Balance   int       `json:"balance,min=1"`
	CreatedAt time.Time `json:"created_at"`
}

type GetUserResponse struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname,min=1"`
	Balance   int       `json:"balance,min=1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
