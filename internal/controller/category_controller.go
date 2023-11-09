package controller

import "github.com/gin-gonic/gin"

type Category interface {
	Create(c *gin.Context)
}
