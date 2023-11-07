package entity

import "time"

type Product struct {
	Id         int
	Title      string
	Price      int
	Stock      int
	CategoryId int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
