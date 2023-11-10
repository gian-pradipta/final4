package controller

import "github.com/gin-gonic/gin"

type Category interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	Update(ctx *gin.Context)
}
