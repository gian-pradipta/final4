package dto

type CreateTransactionRequest struct {
	ProductId int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

type CreateTransactionResponse struct {
	Message         string `json:"message" validate:"required"`
	TransactionBill struct {
		TotalPrice   int    `json:"total_price" validate:"required"`
		Quantity     int    `json:"quantity" validate:"required"`
		ProductTitle string `json:"product_title" validate:"required"`
	} `json:"transaction_bill"`
}

type GetTransactionResponse struct {
	Id         int `json:"id"`
	ProductId  int `json:"product_id"`
	UserId     int `json:"user_id"`
	Quantity   int `json:"quantity"`
	TotalPrice int `json:"total_price"`
	Product    GetProductResponse
}
