package dto

import "time"

type CreateCategoryResponse struct {
	Id                int       `json:"id"`
	Type              string    `json:"type,min=1"`
	SoldProductAmount int       `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
}
