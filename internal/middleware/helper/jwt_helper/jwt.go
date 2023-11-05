package jwthelper

import (
	"errors"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var APPLICATION_NAME = "My Simple JWT App"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

func GenerateJWT(username string, email string) (string, error) {
	var err error
	var signedToken string
	var newClaim UserClaims
	var newToken *jwt.Token
	newClaim.StandardClaims = jwt.StandardClaims{
		Issuer:    APPLICATION_NAME,
		ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
	}
	newClaim.Username = username
	newClaim.Email = email
	newToken = jwt.NewWithClaims(JWT_SIGNING_METHOD, newClaim)
	signedToken, err = newToken.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		log.Fatal(err)
	}
	return signedToken, err
}

func ParseJWT(signedToken string) (UserClaims, error) {
	var claims UserClaims
	var err error
	token, err := jwt.Parse(signedToken, func(t *jwt.Token) (interface{}, error) {

		return JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		return claims, err
	}
	claims, ok := token.Claims.(UserClaims)
	if !ok || !token.Valid {
		err = errors.New("Validation error")
	}
	return claims, err
}
