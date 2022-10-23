package service

import (
	"api-gateway/dto"
	pb "api-gateway/proto"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-micro.dev/v4/errors"
	"io"
	"net/http"
)

type CustomerService interface {
	FindOneByID(ctx context.Context, customerID int) (dto.CustomerResponse, error)
	FindAll(ctx context.Context) ([]dto.CustomerResponse, error)
	Create(ctx context.Context, request dto.CustomerCreateReq) (dto.CustomerResponse, error)
	Update(ctx context.Context, request dto.CustomerUpdateReq) (dto.CustomerResponse, error)
	Delete(ctx context.Context, request dto.CustomerDeleteReq) error
}

func NewCustomerService(
	customerService pb.CustomerService,
) CustomerService {
	return &customerServiceImpl{
		CustomerRPC: customerService,
	}
}

type customerServiceImpl struct {
	CustomerRPC pb.CustomerService
}

func (service *customerServiceImpl) FindOneByID(ctx context.Context, customerID int) (dto.CustomerResponse, error) {
	customer, err := service.CustomerRPC.FindOneByID(ctx, &pb.CustomerID{
		ID: int64(customerID),
	})
	if err != nil {
		e := errors.FromError(err)
		return dto.CustomerResponse{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("customer: %s", e.GetDetail()))
	}
	return dto.CustomerResponse{
		ID:        int(customer.ID),
		Name:      customer.Name,
		Email:     customer.Email,
		Password:  customer.Password,
		CreatedAt: dto.DateTime(customer.CreatedAt.AsTime()),
		UpdatedAt: dto.DateTime(customer.UpdatedAt.AsTime()),
	}, nil
}

func (service *customerServiceImpl) FindAll(ctx context.Context) ([]dto.CustomerResponse, error) {
	customersResponse := make([]dto.CustomerResponse, 0)
	stream, err := service.CustomerRPC.FindAll(ctx, nil)
	if err != nil {
		e := errors.FromError(err)
		return nil, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("customer: %s", e.GetDetail()))
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		customersResponse = append(customersResponse, dto.CustomerResponse{
			ID:        int(msg.ID),
			Name:      msg.Name,
			Email:     msg.Email,
			Password:  msg.Password,
			CreatedAt: dto.DateTime(msg.CreatedAt.AsTime()),
			UpdatedAt: dto.DateTime(msg.UpdatedAt.AsTime()),
		})
	}
	return customersResponse, nil
}

func (service *customerServiceImpl) Create(ctx context.Context, request dto.CustomerCreateReq) (dto.CustomerResponse, error) {
	customer, err := service.CustomerRPC.Create(ctx, &pb.CustomerCreateReq{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		e := errors.FromError(err)
		return dto.CustomerResponse{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("customer: %s", e.GetDetail()))
	}
	return dto.CustomerResponse{
		ID:        int(customer.ID),
		Name:      customer.Name,
		Email:     customer.Email,
		Password:  customer.Password,
		CreatedAt: dto.DateTime(customer.CreatedAt.AsTime()),
		UpdatedAt: dto.DateTime(customer.UpdatedAt.AsTime()),
	}, nil
}

func (service *customerServiceImpl) Update(ctx context.Context, request dto.CustomerUpdateReq) (dto.CustomerResponse, error) {
	customer, err := service.CustomerRPC.Update(ctx, &pb.CustomerUpdateReq{
		ID:       int64(request.ID),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Query: &pb.QueryData{
			ID:   request.QueryData.ID,
			Role: request.QueryData.Role,
		},
	})
	if err != nil {
		e := errors.FromError(err)
		return dto.CustomerResponse{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("customer: %s", e.GetDetail()))
	}
	return dto.CustomerResponse{
		ID:        int(customer.ID),
		Name:      customer.Name,
		Email:     customer.Email,
		Password:  customer.Password,
		CreatedAt: dto.DateTime(customer.CreatedAt.AsTime()),
		UpdatedAt: dto.DateTime(customer.UpdatedAt.AsTime()),
	}, nil
}

func (service *customerServiceImpl) Delete(ctx context.Context, request dto.CustomerDeleteReq) error {
	_, err := service.CustomerRPC.Delete(ctx, &pb.DeleteReq{
		ID: int64(request.ID),
		Query: &pb.QueryData{
			ID:   request.QueryData.ID,
			Role: request.QueryData.Role,
		},
	})
	if err != nil {
		e := errors.FromError(err)
		return echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("customer: %s", e.GetDetail()))
	}
	return nil
}
