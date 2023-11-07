package controller

import (
	"final2/internal/dto"
	"final2/internal/helper/errorhandler"
	jwthelper "final2/internal/helper/jwt_helper"
	"final2/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type user struct {
	s service.User
}

func NewUserController(s service.User) User {
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
	validate := validator.New()
	err = c.ShouldBindJSON(&newUser)
	if err != nil {
		goto ERROR_HANDLING
	}
	err = validate.Struct(&newUser)
	if err != nil {
		goto ERROR_HANDLING
	}
	err = u.s.Login(newUser)
	if err != nil {
		goto ERROR_HANDLING
	}
  token, err = jwthelper.GenerateJWT(newUser.Email)
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
