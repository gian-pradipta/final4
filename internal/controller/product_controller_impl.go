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
	var response dto.GetProductsResponse
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

func (p *product) GetAll(ctx *gin.Context) {
	var response []dto.GetProductsResponse
	var err error
	var errCode int = http.StatusBadRequest
	s := p.serv

	response, err = s.GetAll()
	if err != nil {
		err = errors.New("Failed to get data")
		goto ERROR_HANDLING
	}

ERROR_HANDLING:
	if err != nil {
		httpError := errorhandler.NewHttpError(err.Error(), errCode)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpError)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, response)
}

func (p *product) Update(ctx *gin.Context) {
	var err error
	var errCode int = http.StatusBadRequest
	var request dto.CreateProductRequest
	var response dto.GetProductsResponse
	id, err := getID(ctx)
	if err != nil {
		err = errors.New("Bad Request")
		goto ERROR_HANDLING
	}
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		err = errors.New("Invalid JSON request")
		goto ERROR_HANDLING
	}
	err = p.v.Struct(&request)
	if err != nil {
		err = errors.New("JSON Body violates one or more constraints")
		goto ERROR_HANDLING
	}
	_, err = p.serv.Update(id, request)
	if err != nil {
		err = errors.New("Failed to update data")
		goto ERROR_HANDLING
	}
	response, err = p.serv.Get(id)
	if err != nil {
		err = errors.New("Failed to get data")
		goto ERROR_HANDLING
	}
ERROR_HANDLING:
	if err != nil {
		httpError := errorhandler.NewHttpError(err.Error(), errCode)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpError)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"product": response})
}

func (p *product) Delete(ctx *gin.Context) {
	id, err := getID(ctx)
	errCode := http.StatusBadRequest
	var response dto.OnelineResponse
	response.Message = "Product has been succesfully deleted"
	if err != nil {
		err = errors.New("Bad request")
		goto ERROR_HANDLING
	}
	err = p.serv.Delete(id)
	if err != nil {
		err = errors.New("Failed deleting product: product may not exist")
		goto ERROR_HANDLING
	}
ERROR_HANDLING:
	if err != nil {
		httpError := errorhandler.NewHttpError(err.Error(), errCode)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpError)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, response)
}
