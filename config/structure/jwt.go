package structure

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
