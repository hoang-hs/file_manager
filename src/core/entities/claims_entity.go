package entities

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Id string
	jwt.StandardClaims
}
