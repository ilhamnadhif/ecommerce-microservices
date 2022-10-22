package dto

import (
	pb "api-gateway/proto"
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
	ID       int     `json:"id"`
	Role     pb.Role `json:"role"`
	RoleName string  `json:"role_name"`
	jwt.StandardClaims
}
