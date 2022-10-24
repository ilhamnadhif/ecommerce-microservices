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
	FindOneByID(ctx context.Context, customerID int) (dto.CustomerResponseWithCartProducts, error)
	FindAll(ctx context.Context) ([]dto.CustomerResponse, error)
	Create(ctx context.Context, request dto.CustomerCreateReq) (dto.CustomerResponse, error)
	Update(ctx context.Context, request dto.CustomerUpdateReq) (dto.CustomerResponse, error)
	Delete(ctx context.Context, request dto.CustomerDeleteReq) error
}

func NewCustomerService(
	customerService pb.CustomerService,
	cartService pb.CartService,
	productService pb.ProductService,
) CustomerService {
	return &customerServiceImpl{
		CustomerRPC: customerService,
		CartRPC:     cartService,
		ProductRPC:  productService,
	}
}

type customerServiceImpl struct {
	CustomerRPC pb.CustomerService
	CartRPC     pb.CartService
	ProductRPC  pb.ProductService
}

func (service *customerServiceImpl) FindOneByID(ctx context.Context, customerID int) (dto.CustomerResponseWithCartProducts, error) {
	customer, err := service.CustomerRPC.FindOneByID(ctx, &pb.CustomerID{
		ID: int64(customerID),
	})
	if err != nil {
		e := errors.FromError(err)
		return dto.CustomerResponseWithCartProducts{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("customer: %s", e.GetDetail()))
	}
	cartProductsResponse := make([]dto.CartResponseWithProduct, 0)
	stream, err := service.CartRPC.FindAllByCustomerID(ctx, &pb.CustomerID{ID: int64(customerID)})
	if err != nil {
		e := errors.FromError(err)
		return dto.CustomerResponseWithCartProducts{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("customer: %s", e.GetDetail()))
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return dto.CustomerResponseWithCartProducts{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		product, err := service.ProductRPC.FindOneByID(ctx, &pb.ProductID{
			ID: msg.ProductID,
		})
		if err != nil {
			e := errors.FromError(err)
			return dto.CustomerResponseWithCartProducts{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("customer: %s", e.GetDetail()))
		}
		cartProductsResponse = append(cartProductsResponse, dto.CartResponseWithProduct{
			ID:         int(msg.ID),
			CustomerID: int(msg.CustomerID),
			ProductID:  int(msg.ProductID),
			Quantity:   int(msg.Quantity),
			CreatedAt:  dto.DateTime(msg.CreatedAt.AsTime()),
			UpdatedAt:  dto.DateTime(msg.UpdatedAt.AsTime()),
			Product: dto.ProductResponse{
				ID:          int(product.ID),
				Name:        product.Name,
				MerchantID:  int(product.MerchantID),
				Description: product.Description,
				Price:       int(product.Price),
				CreatedAt:   dto.DateTime(product.CreatedAt.AsTime()),
				UpdatedAt:   dto.DateTime(product.UpdatedAt.AsTime()),
			},
		})
	}
	return dto.CustomerResponseWithCartProducts{
		ID:        int(customer.ID),
		Name:      customer.Name,
		Email:     customer.Email,
		Password:  customer.Password,
		CreatedAt: dto.DateTime(customer.CreatedAt.AsTime()),
		UpdatedAt: dto.DateTime(customer.UpdatedAt.AsTime()),
		Carts:     cartProductsResponse,
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
	if request.QueryData.Role != dto.CUSTOMER_ROLE {
		return dto.CustomerResponse{}, echo.NewHTTPError(http.StatusForbidden, "access denied for this role")
	}
	if request.ID != request.QueryData.ID {
		return dto.CustomerResponse{}, echo.NewHTTPError(http.StatusForbidden, "access denied for this account")
	}
	customer, err := service.CustomerRPC.Update(ctx, &pb.CustomerUpdateReq{
		ID:       int64(request.ID),
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

func (service *customerServiceImpl) Delete(ctx context.Context, request dto.CustomerDeleteReq) error {
	if request.QueryData.Role != dto.CUSTOMER_ROLE {
		return echo.NewHTTPError(http.StatusForbidden, "access denied for this role")
	}
	if request.ID != request.QueryData.ID {
		return echo.NewHTTPError(http.StatusForbidden, "access denied for this account")
	}
	_, err := service.CustomerRPC.Delete(ctx, &pb.CustomerID{
		ID: int64(request.ID),
	})
	if err != nil {
		e := errors.FromError(err)
		return echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("customer: %s", e.GetDetail()))
	}
	return nil
}
