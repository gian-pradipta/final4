package controller

import (
	"errors"
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
	var response dto.CreateUserResponse
	var id int
	err = c.ShouldBindJSON(&newUser)
	validate := validator.New()
	errCode = http.StatusBadRequest
	if err != nil {
		err = errors.New("Invalid JSON Body")
		goto ERROR_HANDLING
	}
	err = validate.Struct(&newUser)
	if err != nil {
		err = errors.New("JSON Body violates one or more constraints")
		goto ERROR_HANDLING
	}
	id, err = s.Create(newUser)
	// fmt.Println("Hello")
	if err != nil {
		err = errors.New("Data Duplication")
		goto ERROR_HANDLING
	}

	response, err = s.GetCreateResponse(id)
	if err != nil {
		err = errors.New("Bad Request")
		goto ERROR_HANDLING
	}

ERROR_HANDLING:
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorhandler.NewHttpError(err.Error(), errCode))
		return
	}
	c.AbortWithStatusJSON(http.StatusCreated, response)
}

func (u *user) Login(c *gin.Context) {
	var err error
	var errCode int = http.StatusBadRequest
	var errMessage string
	var response dto.LoginUserResponse
	var newUser dto.LoginUserRequest
	var token string
	var group string
	validate := validator.New()

	err = c.ShouldBindJSON(&newUser)
	if err != nil {
		errMessage = "Invalid JSON Request"
		goto ERROR_HANDLING
	}
	err = validate.Struct(&newUser)
	if err != nil {
		errMessage = "Invalid JSON Request"
		goto ERROR_HANDLING
	}

	group, err = u.s.Login(newUser)
	if err != nil {
		errMessage = "Login Failed"
		goto ERROR_HANDLING
	}
	token, err = jwthelper.GenerateJWT(newUser.Email, group)
	if err != nil {
		errMessage = "Login Failed"
		goto ERROR_HANDLING
	}
ERROR_HANDLING:
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorhandler.NewHttpError(errMessage, errCode))
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
	_, email, group, err := GetAuthorizedInformation(c)
	if err != nil {
		err = errors.New("Bad request")
		goto ERROR_HANDLING
	}
	err = c.ShouldBindJSON(&user)

	if err != nil {
		err = errors.New("Invalid JSON request")
		errCode = http.StatusBadRequest
		goto ERROR_HANDLING
	}
	err = v.Struct(&user)

	if err != nil {
		err = errors.New("Invalid JSON field value")
		errCode = http.StatusBadRequest
		goto ERROR_HANDLING
	}
	err = s.TopUp(user, email, group)
	if err != nil {
		err = errors.New("Failed to update user's balance")
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
