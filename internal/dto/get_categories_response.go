package dto

import "time"

type GetCategoriesResponse struct {
	Id                int                   `json:"id"`
	Type              string                `json:"type"`
	SoldProductAmount int                   `json:"sold_product_amount"`
	CreatedAt         time.Time             `json:"created_at"`
	UpdatedAt         time.Time             `json:"updated_at"`
	Products          []GetProductsResponse `json:"Products"`
}
