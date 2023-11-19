package controller

import "github.com/gin-gonic/gin"

type Transaction interface {
	Create(ctx *gin.Context)
}
