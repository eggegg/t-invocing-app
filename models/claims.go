package models

import (
	"github.com/dgrijalva/jwt-go"
)

// jwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	Name  string `json:"username"`
	Id  int   `json:"id"`
	jwt.StandardClaims
}