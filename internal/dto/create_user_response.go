package dto

import "time"

type CreateUserResponse struct {
	Id        int       `json:"id"`
	Fullname  string    `json:"fullname,min=1"`
	Email     string    `json:"email"`
	Balance   int       `json:"balance,min=1"`
	CreatedAt time.Time `json:"created_at"`
}
