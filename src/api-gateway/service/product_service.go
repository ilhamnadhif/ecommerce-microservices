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

type ProductService interface {
	FindOneByID(ctx context.Context, productID int) (dto.ProductResponseWithMerchant, error)
	FindAll(ctx context.Context) ([]dto.ProductResponse, error)
	Create(ctx context.Context, request dto.ProductCreateReq) (dto.ProductResponse, error)
	Update(ctx context.Context, request dto.ProductUpdateReq) (dto.ProductResponse, error)
	Delete(ctx context.Context, productID int, queryData dto.QueryData) error
}

func NewProductService(
	productService pb.ProductService,
	merchantService pb.MerchantService,
) ProductService {
	return &productServiceImpl{
		ProductRPC:  productService,
		MerchantRPC: merchantService,
	}
}

type productServiceImpl struct {
	ProductRPC  pb.ProductService
	MerchantRPC pb.MerchantService
}

func (service *productServiceImpl) FindOneByID(ctx context.Context, productID int) (dto.ProductResponseWithMerchant, error) {
	product, err := service.ProductRPC.FindOneByID(ctx, &pb.ProductID{
		ID: int64(productID),
	})
	if err != nil {
		e := errors2.FromError(err)
		return dto.ProductResponseWithMerchant{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("product: %s", e.GetDetail()))
	}
	merchant, err := service.MerchantRPC.FindOneByID(ctx, &pb.MerchantID{
		ID: product.MerchantID,
	})
	if err != nil {
		e := errors2.FromError(err)
		return dto.ProductResponseWithMerchant{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("merchant: %s", e.GetDetail()))
	}
	return dto.ProductResponseWithMerchant{
		ProductResponse: dto.ProductResponse{
			ID:          int(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Price:       int(product.Price),
			CreatedAt:   dto.DateTime(product.CreatedAt.AsTime()),
			UpdatedAt:   dto.DateTime(product.UpdatedAt.AsTime()),
		},
		Merchant: dto.MerchantResponse{
			ID:        int(merchant.ID),
			Name:      merchant.Name,
			Email:     merchant.Email,
			Password:  merchant.Password,
			CreatedAt: dto.DateTime(merchant.CreatedAt.AsTime()),
			UpdatedAt: dto.DateTime(merchant.UpdatedAt.AsTime()),
		},
	}, nil
}

func (service *productServiceImpl) FindAll(ctx context.Context) ([]dto.ProductResponse, error) {
	productsResponse := make([]dto.ProductResponse, 0)
	stream, err := service.ProductRPC.FindAll(ctx, nil)
	if err != nil {
		e := errors2.FromError(err)
		return nil, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("product: %s", e.GetDetail()))
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
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
	return productsResponse, nil
}

func (service *productServiceImpl) Create(ctx context.Context, request dto.ProductCreateReq) (dto.ProductResponse, error) {
	_, err := service.MerchantRPC.FindOneByID(ctx, &pb.MerchantID{
		ID: int64(request.MerchantID),
	})
	if err != nil {
		e := errors2.FromError(err)
		return dto.ProductResponse{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("merchant: %s", e.GetDetail()))
	}
	product, err := service.ProductRPC.Create(ctx, &pb.ProductCreateReq{
		MerchantID:  int64(request.MerchantID),
		Name:        request.Name,
		Description: request.Description,
		Price:       int64(request.Price),
		Query: &pb.QueryData{
			MerchantID: request.QueryData.ID,
			Role:       request.QueryData.Role,
		},
	})
	if err != nil {
		e := errors2.FromError(err)
		return dto.ProductResponse{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("product: %s", e.GetDetail()))
	}
	return dto.ProductResponse{
		ID:          int(product.ID),
		MerchantID:  int(product.MerchantID),
		Name:        product.Name,
		Description: product.Description,
		Price:       int(product.Price),
		CreatedAt:   dto.DateTime(product.CreatedAt.AsTime()),
		UpdatedAt:   dto.DateTime(product.UpdatedAt.AsTime()),
	}, nil
}

func (service *productServiceImpl) Update(ctx context.Context, request dto.ProductUpdateReq) (dto.ProductResponse, error) {
	product, err := service.ProductRPC.Update(ctx, &pb.ProductUpdateReq{
		ID:          int64(request.ID),
		Name:        request.Name,
		Description: request.Description,
		Price:       int64(request.Price),
		Query: &pb.QueryData{
			MerchantID: request.QueryData.ID,
			Role:       request.QueryData.Role,
		},
	})
	if err != nil {
		e := errors2.FromError(err)
		return dto.ProductResponse{}, echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("product: %s", e.GetDetail()))
	}
	return dto.ProductResponse{
		ID:          int(product.ID),
		MerchantID:  int(product.MerchantID),
		Name:        product.Name,
		Description: product.Description,
		Price:       int(product.Price),
		CreatedAt:   dto.DateTime(product.CreatedAt.AsTime()),
		UpdatedAt:   dto.DateTime(product.UpdatedAt.AsTime()),
	}, nil
}

func (service *productServiceImpl) Delete(ctx context.Context, productID int, queryData dto.QueryData) error {
	_, err := service.ProductRPC.Delete(ctx, &pb.DeleteReq{
		ID: int64(productID),
		Query: &pb.QueryData{
			MerchantID: queryData.ID,
			Role:       queryData.Role,
		},
	})
	if err != nil {
		e := errors2.FromError(err)
		return echo.NewHTTPError(int(e.GetCode()), fmt.Sprintf("product: %s", e.GetDetail()))
	}
	return nil
}
