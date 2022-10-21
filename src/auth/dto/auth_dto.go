package dto

import "github.com/golang-jwt/jwt"

type JWTCustomClaims struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}
