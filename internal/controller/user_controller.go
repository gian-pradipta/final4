package controller

import "github.com/gin-gonic/gin"

type User interface {
	Create(c *gin.Context)
	Login(c *gin.Context)
}
