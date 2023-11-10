package entity

import "time"

type CategoryWithProduct struct {
	Id                int
	Type              string
	SoldProductAmount int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Products          []Product
}
