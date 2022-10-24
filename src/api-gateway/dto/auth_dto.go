package dto

import (
	"github.com/golang-jwt/jwt"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type JWTCustomClaims struct {
	ID   int  `json:"id"`
	Role Role `json:"role"`
	jwt.StandardClaims
}
