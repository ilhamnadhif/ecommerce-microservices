package dto

import (
	"github.com/golang-jwt/jwt"
	"net/http"
)

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func WebResponseSuccess(data interface{}) WebResponse {
	return WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   data,
	}
}

type JWTCustomClaims struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
