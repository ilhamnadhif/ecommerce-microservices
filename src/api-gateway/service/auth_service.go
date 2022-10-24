package service

import (
	"api-gateway/dto"
	"api-gateway/pkg"
	pb "api-gateway/proto"
	"context"
	"github.com/labstack/echo/v4"
	errors2 "go-micro.dev/v4/errors"
	"net/http"
)

type AuthService interface {
	LoginMerchant(ctx context.Context, req dto.LoginReq) (dto.TokenResponse, error)
	LoginCustomer(ctx context.Context, req dto.LoginReq) (dto.TokenResponse, error)
}

func NewAuthService(
	merchantService pb.MerchantService,
	customerService pb.CustomerService,
) AuthService {
	return &authServiceImpl{
		MerchantRPC: merchantService,
		CustomerRPC: customerService,
	}
}

type authServiceImpl struct {
	MerchantRPC pb.MerchantService
	CustomerRPC pb.CustomerService
}

func (service *authServiceImpl) LoginMerchant(ctx context.Context, req dto.LoginReq) (dto.TokenResponse, error) {
	merchant, err := service.MerchantRPC.FindOneByEmail(ctx, &pb.MerchantEmail{
		Email: req.Email,
	})
	if err != nil {
		e := errors2.FromError(err)
		return dto.TokenResponse{}, echo.NewHTTPError(int(e.GetCode()), e.GetDetail())
	}
	if req.Password != merchant.Password {
		return dto.TokenResponse{}, echo.NewHTTPError(http.StatusBadRequest, "password not match")
	}
	token, err := pkg.GenerateToken(dto.JWTCustomClaims{
		ID:   int(merchant.ID),
		Role: dto.MERCHANT_ROLE,
	})
	if err != nil {
		return dto.TokenResponse{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return dto.TokenResponse{
		Token: token,
	}, nil
}

func (service *authServiceImpl) LoginCustomer(ctx context.Context, req dto.LoginReq) (dto.TokenResponse, error) {
	customer, err := service.CustomerRPC.FindOneByEmail(ctx, &pb.CustomerEmail{
		Email: req.Email,
	})
	if err != nil {
		e := errors2.FromError(err)
		return dto.TokenResponse{}, echo.NewHTTPError(int(e.GetCode()), e.GetDetail())
	}
	if req.Password != customer.Password {
		return dto.TokenResponse{}, echo.NewHTTPError(http.StatusBadRequest, "password not match")
	}
	token, err := pkg.GenerateToken(dto.JWTCustomClaims{
		ID:   int(customer.ID),
		Role: dto.CUSTOMER_ROLE,
	})
	if err != nil {
		return dto.TokenResponse{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return dto.TokenResponse{
		Token: token,
	}, nil
}
