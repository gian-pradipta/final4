package jwthelper

import jwt "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	jwt.StandardClaims
	Email string `json:"Email"`
  Group string `json:"group"`
}
