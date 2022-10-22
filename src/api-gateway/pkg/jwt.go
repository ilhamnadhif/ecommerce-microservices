package pkg

import (
	"api-gateway/config"
	"api-gateway/dto"
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateToken(claims dto.JWTCustomClaims) (string, error) {
	claims.ExpiresAt = time.Now().Add(time.Hour * time.Duration(config.Config.Jwt.ExpiredIn)).Unix()
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.Config.Jwt.SigningKey))
	if err != nil {
		return "", err
	}
	return t, nil
}
