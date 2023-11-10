package controller

import (
	"final2/internal/dto"
	"final2/internal/helper/errorhandler"
	"final2/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type category struct {
	serv service.Category
	v    *validator.Validate
}

func NewCategory(s service.Category, v *validator.Validate) Category {
	var c category
	c.serv = s
	c.v = v
	return &c
}

func (c *category) Create(ctx *gin.Context) {
	s := c.serv
	var errCode int = http.StatusBadRequest
	var err error
	var latestId int
	var newCategory dto.CreateCategoryRequest
	var response dto.CreateCategoryResponse
	err = ctx.ShouldBindJSON(&newCategory)
	if err != nil {
		goto ERROR_HANDLING
	}
	latestId, err = s.Create(newCategory)
	if err != nil {
		goto ERROR_HANDLING
	}
	response, err = s.Get(latestId)
	if err != nil {
		goto ERROR_HANDLING
	}

ERROR_HANDLING:
	if err != nil {
		httpError := errorhandler.NewHttpError(err.Error(), errCode)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpError)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusCreated, response)
}
