package controller

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type category struct {
	db *sql.DB
}

func NewCategory(db *sql.DB) Category {
	var c category
	c.db = db
	return &c
}

func (c *category) Create(ctx *gin.Context) {
	// db :=
}
