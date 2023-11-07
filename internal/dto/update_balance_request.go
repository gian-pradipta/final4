package dto

type UpdateBalanceRequest struct {
	Balance int `json:"balance" validate:"required,min=0,max=100000000"`
}
