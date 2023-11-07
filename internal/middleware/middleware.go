package middleware

import (
	"final2/internal/helper/errorhandler"
	jwthelper "final2/internal/helper/jwt_helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {

			c.AbortWithStatusJSON(http.StatusUnauthorized, errorhandler.NewHttpError("unauthorized access", http.StatusUnauthorized))
			return
		}
		t := strings.Split(authHeader, " ")
		var authToken string
		if len(t) == 2 {
			authToken = t[1]
		}

		data, err := jwthelper.ParseJWT(authToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorhandler.NewHttpError(err.Error(), http.StatusUnauthorized))
			return
		}
		c.Set("group", data.Group)
		c.Next()
	}
}

func Authorize() gin.HandlerFunc {

	return func(c *gin.Context) {
		group, exist := c.Get("group")
		if !exist || group != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorhandler.NewHttpError("unauthorized access", http.StatusUnauthorized))
			return
		}

		c.Next()
	}
}
