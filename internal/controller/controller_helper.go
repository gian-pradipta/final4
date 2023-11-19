package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAuthorizedInformation(c *gin.Context) (int, string, string, error) {
	var err error
	var email string
	var group string
	var id int
	eml, exist := c.Get("email")
	if !exist {
		err = errors.New("Authorization Failed")
		return id, email, group, err
	}

	grp, exist := c.Get("group")
	if !exist {
		err = errors.New("Authorization Failed")
		return id, email, group, err
	}
	userId, exist := c.Get("userId")
	if !exist {
		err = errors.New("Authorization Failed")
		return id, email, group, err
	}
	stringId := userId.(string)

	id, _ = strconv.Atoi(stringId)
	email = eml.(string)
	group = grp.(string)

	return id, email, group, err
}

func getID(c *gin.Context) (int, error) {
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	return id, err
}
