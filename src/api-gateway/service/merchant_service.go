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

type MerchantService interface {
	FindOneByID(ctx context.Context, merchantID int) (dto.MerchantResponseWithProducts, error)
	FindAll(ctx context.Context) ([]dto.MerchantResponse, error)
	Create(ctx context.Context, request dto.MerchantCreateReq) (dto.MerchantResponse, error)
	Update(ctx context.Context, request dto.MerchantUpdateReq) (dto.MerchantResponse, error)
	Delete(ctx context.Context, merchantID int, queryData dto.QueryData) error
}

func NewMerchantService(
	merchantService pb.MerchantService,
	productService pb.ProductService,
) MerchantService {
	return &merchantServiceImpl{
		MerchantRPC: merchantService,
		ProductRPC:  productService,
	}
}

type merchantServiceImpl struct {
	MerchantRPC pb.MerchantService
	ProductRPC  pb.ProductService
}

func (service *merchantServiceImpl) FindOneByID(ctx context.Context, merchantID int) (dto.MerchantResponseWithProducts, error) {
	merchant, err := service.MerchantRPC.FindOneByID(ctx, &pb.MerchantID{
		ID: int64(merchantID),
	})
	if err != nil {
		e := errors.FromError(err)
		return dto.MerchantResponseWithProducts{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("merchant: %s", e.GetDetail()))
	}
	stream, err := service.ProductRPC.FindAllByMerchantID(ctx, &pb.MerchantID{ID: int64(merchantID)})
	if err != nil {
		e := errors.FromError(err)
		return dto.MerchantResponseWithProducts{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("product: %s", e.GetDetail()))
	}
	productsResponse := make([]dto.ProductResponse, 0)
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return dto.MerchantResponseWithProducts{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		productsResponse = append(productsResponse, dto.ProductResponse{
			ID:          int(msg.ID),
			MerchantID:  int(msg.MerchantID),
			Name:        msg.Name,
			Description: msg.Description,
			Price:       int(msg.Price),
			CreatedAt:   dto.DateTime(msg.CreatedAt.AsTime()),
			UpdatedAt:   dto.DateTime(msg.UpdatedAt.AsTime()),
		})
	}
	return dto.MerchantResponseWithProducts{
		MerchantResponse: dto.MerchantResponse{
			ID:        int(merchant.ID),
			Name:      merchant.Name,
			Email:     merchant.Email,
			Password:  merchant.Password,
			CreatedAt: dto.DateTime(merchant.CreatedAt.AsTime()),
			UpdatedAt: dto.DateTime(merchant.UpdatedAt.AsTime()),
		},
		Products: productsResponse,
	}, nil
}

func (service *merchantServiceImpl) FindAll(ctx context.Context) ([]dto.MerchantResponse, error) {
	merchantsResponse := make([]dto.MerchantResponse, 0)
	stream, err := service.MerchantRPC.FindAll(ctx, nil)
	if err != nil {
		e := errors.FromError(err)
		return nil, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("merchant: %s", e.GetDetail()))
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		merchantsResponse = append(merchantsResponse, dto.MerchantResponse{
			ID:        int(msg.ID),
			Name:      msg.Name,
			Email:     msg.Email,
			Password:  msg.Password,
			CreatedAt: dto.DateTime(msg.CreatedAt.AsTime()),
			UpdatedAt: dto.DateTime(msg.UpdatedAt.AsTime()),
		})
	}
	return merchantsResponse, nil
}

func (service *merchantServiceImpl) Create(ctx context.Context, request dto.MerchantCreateReq) (dto.MerchantResponse, error) {
	merchant, err := service.MerchantRPC.Create(ctx, &pb.MerchantCreateReq{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		e := errors.FromError(err)
		return dto.MerchantResponse{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("merchant: %s", e.GetDetail()))
	}
	return dto.MerchantResponse{
		ID:        int(merchant.ID),
		Name:      merchant.Name,
		Email:     merchant.Email,
		Password:  merchant.Password,
		CreatedAt: dto.DateTime(merchant.CreatedAt.AsTime()),
		UpdatedAt: dto.DateTime(merchant.UpdatedAt.AsTime()),
	}, nil
}

func (service *merchantServiceImpl) Update(ctx context.Context, request dto.MerchantUpdateReq) (dto.MerchantResponse, error) {
	merchant, err := service.MerchantRPC.Update(ctx, &pb.MerchantUpdateReq{
		ID:       int64(request.ID),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Query: &pb.QueryData{
			MerchantID: request.QueryData.ID,
			Role:       request.QueryData.Role,
		},
	})
	if err != nil {
		e := errors.FromError(err)
		return dto.MerchantResponse{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("merchant: %s", e.GetDetail()))
	}
	return dto.MerchantResponse{
		ID:        int(merchant.ID),
		Name:      merchant.Name,
		Email:     merchant.Email,
		Password:  merchant.Password,
		CreatedAt: dto.DateTime(merchant.CreatedAt.AsTime()),
		UpdatedAt: dto.DateTime(merchant.UpdatedAt.AsTime()),
	}, nil
}

func (service *merchantServiceImpl) Delete(ctx context.Context, merchantID int, queryData dto.QueryData) error {
	stream, err := service.ProductRPC.FindAllByMerchantID(ctx, &pb.MerchantID{ID: int64(merchantID)})
	if err != nil {
		e := errors.FromError(err)
		return echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("product: %s", e.GetDetail()))
	}
	productsResponse := make([]dto.ProductResponse, 0)
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		productsResponse = append(productsResponse, dto.ProductResponse{
			ID:          int(msg.ID),
			MerchantID:  int(msg.MerchantID),
			Name:        msg.Name,
			Description: msg.Description,
			Price:       int(msg.Price),
			CreatedAt:   dto.DateTime(msg.CreatedAt.AsTime()),
			UpdatedAt:   dto.DateTime(msg.UpdatedAt.AsTime()),
		})
	}
	if len(productsResponse) > 0 {
		return echo.NewHTTPError(http.StatusConflict, "products in this merchant is exist")
	}
	_, err = service.MerchantRPC.Delete(ctx, &pb.DeleteReq{
		ID: int64(merchantID),
		Query: &pb.QueryData{
			MerchantID: queryData.ID,
			Role:       queryData.Role,
		},
	})
	if err != nil {
		e := errors.FromError(err)
		return echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("merchant: %s", e.GetDetail()))
	}
	return nil
}
