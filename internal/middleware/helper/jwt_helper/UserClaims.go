package jwthelper

import jwt "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"Username"`
	Email    string `json:"Email"`
}
