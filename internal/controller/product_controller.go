package controller

import "github.com/gin-gonic/gin"

type Product interface {
	Create(ctx *gin.Context)
}
