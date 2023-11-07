package dto

type CreateCategoryRequest struct {
	Type string `json:"type" validate:"required"`
}
