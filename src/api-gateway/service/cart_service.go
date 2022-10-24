package service

import (
	"api-gateway/dto"
	pb "api-gateway/proto"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	errors2 "go-micro.dev/v4/errors"
	"io"
	"net/http"
)

type CartService interface {
	FindOneByID(ctx context.Context, cartID int) (dto.CartResponse, error)
	FindAll(ctx context.Context) ([]dto.CartResponse, error)
	Create(ctx context.Context, request dto.CartCreateReq) (dto.CartResponse, error)
	Update(ctx context.Context, request dto.CartUpdateReq) (dto.CartResponse, error)
	Delete(ctx context.Context, request dto.CartDeleteReq) error
}

func NewCartService(
	cartService pb.CartService,
	productService pb.ProductService,
) CartService {
	return &cartServiceImpl{
		CartRPC:    cartService,
		ProductRPC: productService,
	}
}

type cartServiceImpl struct {
	CartRPC    pb.CartService
	ProductRPC pb.ProductService
}

func (service *cartServiceImpl) FindOneByID(ctx context.Context, cartID int) (dto.CartResponse, error) {
	cart, err := service.CartRPC.FindOneByID(ctx, &pb.CartID{
		ID: int64(cartID),
	})
	if err != nil {
		e := errors2.FromError(err)
		return dto.CartResponse{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("cart: %s", e.GetDetail()))
	}
	return dto.CartResponse{
		ID:         int(cart.ID),
		CustomerID: int(cart.CustomerID),
		ProductID:  int(cart.ProductID),
		Quantity:   int(cart.Quantity),
		CreatedAt:  dto.DateTime(cart.CreatedAt.AsTime()),
		UpdatedAt:  dto.DateTime(cart.UpdatedAt.AsTime()),
	}, nil
}

func (service *cartServiceImpl) FindAll(ctx context.Context) ([]dto.CartResponse, error) {
	cartsResponse := make([]dto.CartResponse, 0)
	stream, err := service.CartRPC.FindAll(ctx, nil)
	if err != nil {
		e := errors2.FromError(err)
		return nil, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("cart: %s", e.GetDetail()))
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		cartsResponse = append(cartsResponse, dto.CartResponse{
			ID:         int(msg.ID),
			CustomerID: int(msg.CustomerID),
			ProductID:  int(msg.ProductID),
			Quantity:   int(msg.Quantity),
			CreatedAt:  dto.DateTime(msg.CreatedAt.AsTime()),
			UpdatedAt:  dto.DateTime(msg.UpdatedAt.AsTime()),
		})
	}
	return cartsResponse, nil
}

func (service *cartServiceImpl) Create(ctx context.Context, request dto.CartCreateReq) (dto.CartResponse, error) {
	if request.QueryData.Role != dto.CUSTOMER_ROLE {
		return dto.CartResponse{}, echo.NewHTTPError(http.StatusForbidden, "access denied for this role")
	}
	_, err := service.ProductRPC.FindOneByID(ctx, &pb.ProductID{
		ID: int64(request.ProductID),
	})
	if err != nil {
		e := errors2.FromError(err)
		return dto.CartResponse{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("product: %s", e.GetDetail()))
	}
	cart, err := service.CartRPC.Create(ctx, &pb.CartCreateReq{
		CustomerID: int64(request.QueryData.ID),
		ProductID:  int64(request.ProductID),
		Quantity:   int64(request.Quantity),
	})
	if err != nil {
		e := errors2.FromError(err)
		return dto.CartResponse{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("cart: %s", e.GetDetail()))
	}
	return dto.CartResponse{
		ID:         int(cart.ID),
		CustomerID: int(cart.CustomerID),
		ProductID:  int(cart.ProductID),
		Quantity:   int(cart.Quantity),
		CreatedAt:  dto.DateTime(cart.CreatedAt.AsTime()),
		UpdatedAt:  dto.DateTime(cart.UpdatedAt.AsTime()),
	}, nil
}

func (service *cartServiceImpl) Update(ctx context.Context, request dto.CartUpdateReq) (dto.CartResponse, error) {
	if request.QueryData.Role != dto.CUSTOMER_ROLE {
		return dto.CartResponse{}, echo.NewHTTPError(http.StatusForbidden, "access denied for this role")
	}
	findCart, err := service.CartRPC.FindOneByID(ctx, &pb.CartID{
		ID: int64(request.ID),
	})
	if err != nil {
		e := errors2.FromError(err)
		return dto.CartResponse{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("cart: %s", e.GetDetail()))
	}
	if request.QueryData.ID != int(findCart.CustomerID) {
		return dto.CartResponse{}, echo.NewHTTPError(http.StatusForbidden, "access denied for this account")
	}
	cart, err := service.CartRPC.Update(ctx, &pb.CartUpdateReq{
		ID:       int64(request.ID),
		Quantity: int64(request.Quantity),
	})
	if err != nil {
		e := errors2.FromError(err)
		return dto.CartResponse{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("cart: %s", e.GetDetail()))
	}
	return dto.CartResponse{
		ID:         int(cart.ID),
		CustomerID: int(cart.CustomerID),
		ProductID:  int(cart.ProductID),
		Quantity:   int(cart.Quantity),
		CreatedAt:  dto.DateTime(cart.CreatedAt.AsTime()),
		UpdatedAt:  dto.DateTime(cart.UpdatedAt.AsTime()),
	}, nil
}

func (service *cartServiceImpl) Delete(ctx context.Context, request dto.CartDeleteReq) error {
	if request.QueryData.Role != dto.CUSTOMER_ROLE {
		return echo.NewHTTPError(http.StatusForbidden, "access denied for this role")
	}
	findCart, err := service.CartRPC.FindOneByID(ctx, &pb.CartID{
		ID: int64(request.ID),
	})
	if err != nil {
		e := errors2.FromError(err)
		return echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("cart: %s", e.GetDetail()))
	}
	if request.QueryData.ID != int(findCart.CustomerID) {
		return echo.NewHTTPError(http.StatusForbidden, "access denied for this account")
	}
	_, err = service.CartRPC.Delete(ctx, &pb.CartID{
		ID: int64(request.ID),
	})
	if err != nil {
		e := errors2.FromError(err)
		return echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("cart: %s", e.GetDetail()))
	}
	return nil
}
