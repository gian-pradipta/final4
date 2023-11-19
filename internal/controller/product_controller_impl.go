package controller

import (
	"errors"
	"final2/internal/dto"
	"final2/internal/helper/errorhandler"
	"final2/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type product struct {
	serv service.Product
	v    *validator.Validate
}

func NewProduct(s service.Product, v *validator.Validate) Product {
	var p product
	p.serv = s
	p.v = v
	return &p
}

func (p *product) Create(ctx *gin.Context) {
	var err error
	var errCode int = http.StatusBadRequest
	var response dto.GetProductResponse
	var newProduct dto.CreateProductRequest
	var id int
	s := p.serv
	err = ctx.ShouldBindJSON(&newProduct)
	if err != nil {
		err = errors.New("Invalid JSON Request")
		goto ERROR_HANDLING
	}

	err = p.v.Struct(&newProduct)
	if err != nil {
		err = errors.New("JSON Body violates one or more constraints")
		goto ERROR_HANDLING
	}
	id, err = s.Create(newProduct)
	if err != nil {
		err = errors.New("Failed to Create new product")
		goto ERROR_HANDLING
	}

	response, err = s.Get(id)
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
