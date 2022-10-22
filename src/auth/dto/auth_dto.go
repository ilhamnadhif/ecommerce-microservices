package dto

import (
	pb "auth/proto"
	"github.com/golang-jwt/jwt"
)

type JWTCustomClaims struct {
	ID       int     `json:"id"`
	Role     pb.Role `json:"role"`
	RoleName string  `json:"role_name"`
	jwt.StandardClaims
}
