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

type transaction struct {
	serv service.Transaction
	v    *validator.Validate
}

func NewTransaction(serv service.Transaction, v *validator.Validate) Transaction {
	var t transaction
	t.serv = serv
	t.v = v
	return &t
}

func (t *transaction) Create(ctx *gin.Context) {
	var err error
	var errCode int = http.StatusBadRequest
	var request dto.CreateTransactionRequest
	var response dto.CreateTransactionResponse
	_, email, _, err := GetAuthorizedInformation(ctx)
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		err = errors.New("Invalid JSON body")
		goto ERROR_HANDLING
	}

	err = t.v.Struct(&request)
	if err != nil {
		err = errors.New("Invalid JSON field value")
		goto ERROR_HANDLING
	}

	response, err = t.serv.Create(request, email)
ERROR_HANDLING:
	if err != nil {
		httpError := errorhandler.NewHttpError(err.Error(), errCode)
		ctx.AbortWithStatusJSON(errCode, httpError)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusCreated, response)

}

func (t *transaction) GetMyTransactions(ctx *gin.Context) {
	var err error
	var errCode int = http.StatusBadRequest
	var response []dto.GetTransactionResponse = make([]dto.GetTransactionResponse, 0)

	_, email, _, err := GetAuthorizedInformation(ctx)
	if err != nil {
		goto ERROR_HANDLING
	}
	response, err = t.serv.GetMyTransactions(email)
	if err != nil {
		err = errors.New("User Not Found")
		goto ERROR_HANDLING
	}
ERROR_HANDLING:
	if err != nil {
		httpError := errorhandler.NewHttpError(err.Error(), errCode)
		ctx.AbortWithStatusJSON(errCode, httpError)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, response)
}

func (t *transaction) GetAllTransactions(ctx *gin.Context) {
	var err error
	var errCode int = http.StatusBadRequest
	var response []dto.GetAllTransactionResponse

	response, err = t.serv.GetAllTransactions()
	if err != nil {
		err = errors.New("Bad request")
		goto ERROR_HANDLING
	}

ERROR_HANDLING:
	if err != nil {
		httpError := errorhandler.NewHttpError(err.Error(), errCode)
		ctx.AbortWithStatusJSON(errCode, httpError)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, response)
}
