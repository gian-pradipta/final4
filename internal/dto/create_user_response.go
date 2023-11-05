package dto

import "time"

type CreateUserResponse struct {
	Id        int       `json:"id"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Balance   string    `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}
