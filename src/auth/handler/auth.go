package handler

import (
	"auth/dto"
	"auth/pkg"
	pb "auth/proto"
	"context"
	errors2 "go-micro.dev/v4/errors"
)

func NewAuthHandler(
	merchantService pb.MerchantService,
	customerService pb.CustomerService,
) AuthServiceHandler {
	return AuthServiceHandler{
		MerchantRPC: merchantService,
		CustomerRPC: customerService,
	}
}

type AuthServiceHandler struct {
	MerchantRPC pb.MerchantService
	CustomerRPC pb.CustomerService
}

func (service *AuthServiceHandler) LoginMerchant(ctx context.Context, req *pb.LoginReq, response *pb.TokenResponse) error {
	merchant, err := service.MerchantRPC.FindOneByEmail(ctx, &pb.MerchantEmail{
		Email: req.Email,
	})
	if err != nil {
		return err
	}
	if req.Password != merchant.Password {
		return errors2.Unauthorized("", "password not match")
	}

	token, err := pkg.GenerateToken(dto.JWTCustomClaims{
		ID:       int(merchant.ID),
		Role:     pb.Role_MERCHANT,
		RoleName: "merchant",
	})
	if err != nil {
		return errors2.BadRequest("", err.Error())
	}
	*response = pb.TokenResponse{
		Token: token,
	}
	return nil

}

func (service *AuthServiceHandler) LoginCustomer(ctx context.Context, req *pb.LoginReq, response *pb.TokenResponse) error {
	customer, err := service.CustomerRPC.FindOneByEmail(ctx, &pb.CustomerEmail{
		Email: req.Email,
	})
	if err != nil {
		return err
	}
	if req.Password != customer.Password {
		return errors2.Unauthorized("", "password not match")
	}

	token, err := pkg.GenerateToken(dto.JWTCustomClaims{
		ID:       int(customer.ID),
		Role:     pb.Role_CUSTOMER,
		RoleName: "customer",
	})
	if err != nil {
		return errors2.BadRequest("", err.Error())
	}
	*response = pb.TokenResponse{
		Token: token,
	}
	return nil
}
