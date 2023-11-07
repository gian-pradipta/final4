package entity

import "time"

type Comment struct {
	id         int
	ProductId  int
	UserId     int
	Quantity   int
	TotalPrice int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
