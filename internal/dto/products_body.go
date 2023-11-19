package dto

import "time"

type GetProductsResponse struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryId int       `json:"category_id" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	Title      string `json:"title" validate:"required,min=1"`
	Price      int    `json:"price" validate:"required,min=0,max=50000000"`
	Stock      int    `json:"stock" validate:"required,min=5"`
	CategoryId int    `json:"category_id" validate:"required"`
}

type CreateProductResponse struct {
	Id         int       `json:"id"`
	Title      string    `json:"title" validate:"required"`
	Price      int       `json:"price" validate:"required"`
	Stock      int       `json:"stock" validate:"required,min=5"`
	CategoryId int       `json:"category_id" validate:"required"`
	CreatedAt  time.Time `json:"created_at" validate:"required"`
}

type GetProductResponse struct {
	Id         int       `json:"id"`
	Title      string    `json:"title" validate:"required"`
	Price      int       `json:"price" validate:"required"`
	Stock      int       `json:"stock" validate:"required,min=5"`
	CategoryId int       `json:"category_id" validate:"required"`
	CreatedAt  time.Time `json:"created_at" validate:"required"`
	UpdatedAt  time.Time `json:"updated_at" validate:"required"`
}
