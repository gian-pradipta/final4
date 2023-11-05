package entity

import "time"

type Comment struct {
	id        int
	PhotoId   int
	UserId    int
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
