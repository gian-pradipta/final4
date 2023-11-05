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
		t := strings.Split(authHeader, " ")
		var authToken string
		if len(t) == 2 {
			authToken = t[1]
		}

		_, err := jwthelper.ParseJWT(authToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorhandler.NewHttpError(err.Error(), http.StatusUnauthorized))
			return
		}
		c.Next()
	}
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
