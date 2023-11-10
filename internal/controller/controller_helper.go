package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAuthorizedInformation(c *gin.Context) (string, string, error) {
	var err error
	var email string
	var group string
	eml, exist := c.Get("email")
	if !exist {
		err = errors.New("Authorization Failed")
		return email, group, err
	}

	grp, exist := c.Get("group")
	if !exist {
		err = errors.New("Authorization Failed")
		return email, group, err
	}
	email = eml.(string)
	group = grp.(string)
	return email, group, err
}

func getID(c *gin.Context) (int, error) {
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	return id, err
}
