package entity

type TransactionWithProduct struct {
	Id         int
	ProductId  int
	UserId     int
	Quantity   int
	TotalPrice int
	Product    Product
}

type TransactionWithProductUser struct {
	Id         int
	ProductId  int
	UserId     int
	Quantity   int
	TotalPrice int
	Product    Product
	User       User
}
