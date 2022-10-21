package service

import (
	"api-gateway/dto"
	pb "api-gateway/proto"
	"context"
	"github.com/labstack/echo/v4"
	errors2 "go-micro.dev/v4/errors"
)

type AuthService interface {
	LoginMerchant(ctx context.Context, req dto.LoginReq) (dto.TokenResponse, error)
	LoginCustomer(ctx context.Context, req dto.LoginReq) (dto.TokenResponse, error)
}

func NewAuthService(
	authService pb.AuthService,
) AuthService {
	return &authServiceImpl{
		AuthRPC: authService,
	}
}

type authServiceImpl struct {
	AuthRPC pb.AuthService
}

func (service *authServiceImpl) LoginMerchant(ctx context.Context, req dto.LoginReq) (dto.TokenResponse, error) {
	token, err := service.AuthRPC.LoginMerchant(ctx, &pb.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		e := errors2.FromError(err)
		return dto.TokenResponse{}, echo.NewHTTPError(int(e.GetCode()), e.GetDetail())
	}
	return dto.TokenResponse{
		Token: token.Token,
	}, nil
}

func (service *authServiceImpl) LoginCustomer(ctx context.Context, req dto.LoginReq) (dto.TokenResponse, error) {
	token, err := service.AuthRPC.LoginCustomer(ctx, &pb.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		e := errors2.FromError(err)
		return dto.TokenResponse{}, echo.NewHTTPError(int(e.GetCode()), e.GetDetail())
	}
	return dto.TokenResponse{
		Token: token.Token,
	}, nil
}
