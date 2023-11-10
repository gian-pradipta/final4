package controller

import (
	"final2/internal/dto"
	"final2/internal/helper/errorhandler"
	jwthelper "final2/internal/helper/jwt_helper"
	"final2/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type user struct {
	s service.User
	v *validator.Validate
}

func NewUserController(s service.User, v *validator.Validate) User {
	var u user
	u.s = s
	return &u
}

func (u *user) Create(c *gin.Context) {
	s := u.s
	var err error
	var errCode int
	var newUser dto.CreateUserRequest
	err = c.ShouldBindJSON(&newUser)
	validate := validator.New()
	errCode = http.StatusBadRequest
	if err != nil {
		goto ERROR_HANDLING
	}
	err = validate.Struct(&newUser)
	if err != nil {
		goto ERROR_HANDLING
	}

	err = s.Create(newUser)
	if err != nil {
		goto ERROR_HANDLING
	}

ERROR_HANDLING:
	if err != nil {
		httpError := errorhandler.NewHttpError(err.Error(), errCode)
		c.AbortWithStatusJSON(http.StatusBadRequest, httpError)
		return
	}
	c.AbortWithStatus(http.StatusCreated)
}

func (u *user) Login(c *gin.Context) {
	var err error
	var errCode int = http.StatusBadRequest
	var response dto.LoginUserResponse
	var newUser dto.LoginUserRequest
	var token string
	var group string
	validate := validator.New()

	err = c.ShouldBindJSON(&newUser)
	if err != nil {
		goto ERROR_HANDLING
	}
	err = validate.Struct(&newUser)
	if err != nil {
		goto ERROR_HANDLING
	}

	group, err = u.s.Login(newUser)
	if err != nil {
		goto ERROR_HANDLING
	}
	token, err = jwthelper.GenerateJWT(newUser.Email, group)
	if err != nil {
		goto ERROR_HANDLING
	}
ERROR_HANDLING:
	if err != nil {
		var httpError dto.HttpError
		httpError.Code = errCode
		httpError.Err = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, httpError)
		return
	}
	response.Token = token
	c.AbortWithStatusJSON(http.StatusOK, response)
}

func (u *user) TopUp(c *gin.Context) {
	var err error
	var errCode int = http.StatusUnauthorized
	var user dto.UpdateBalanceRequest
	var response dto.OnelineResponse

	s := u.s
	v := validator.New()
	email, group, err := GetAuthorizedInformation(c)
	if err != nil {
		goto ERROR_HANDLING
	}
	err = c.ShouldBindJSON(&user)

	if err != nil {
		errCode = http.StatusBadRequest
		goto ERROR_HANDLING
	}
	err = v.Struct(&user)

	if err != nil {
		errCode = http.StatusBadRequest
		goto ERROR_HANDLING
	}
	err = s.TopUp(user, email, group)
	if err != nil {
		errCode = http.StatusBadRequest
		goto ERROR_HANDLING
	}
	response.Message = fmt.Sprintf("Your Balance Has Been Successfully Updated to %d", user.Balance)
ERROR_HANDLING:
	if err != nil {
		var httpError dto.HttpError
		httpError.Code = errCode
		httpError.Err = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, httpError)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, response)
}
