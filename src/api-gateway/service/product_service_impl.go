package service

import (
	"api-gateway/dto"
	pb "api-gateway/proto"
	"context"
	"github.com/labstack/echo/v4"
	"go-micro.dev/v4/errors"
	"io"
	"net/http"
)

func NewProductService(service pb.ProductService) ProductService {
	return &productServiceImpl{
		ProductRPC: service,
	}
}

type productServiceImpl struct {
	ProductRPC pb.ProductService
}

func (service *productServiceImpl) FindOneByID(ctx context.Context, productID int) (dto.ProductResponse, error) {
	product, err := service.ProductRPC.FindOneByID(ctx, &pb.ProductID{
		ID: int64(productID),
	})
	if err != nil {
		e := errors.FromError(err)
		return dto.ProductResponse{}, echo.NewHTTPError(int(e.GetCode()), e.GetDetail())
	}
	return dto.ProductResponse{
		ID:          int(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       int(product.Price),
		CreatedAt:   dto.DateTime(product.CreatedAt.AsTime()),
		UpdatedAt:   dto.DateTime(product.UpdatedAt.AsTime()),
	}, nil
}

func (service *productServiceImpl) FindAll(ctx context.Context) ([]dto.ProductResponse, error) {
	stream, err := service.ProductRPC.FindAll(ctx, nil)
	if err != nil {
		e := errors.FromError(err)
		return nil, echo.NewHTTPError(int(e.GetCode()), e.GetDetail())
	}
	productsResponse := make([]dto.ProductResponse, 0)
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
	product, err := service.ProductRPC.Create(ctx, &pb.ProductCreateReq{
		Name:        request.Name,
		Description: request.Description,
		Price:       int64(request.Price),
	})
	if err != nil {
		e := errors.FromError(err)
		return dto.ProductResponse{}, echo.NewHTTPError(int(e.GetCode()), e.GetDetail())
	}
	return dto.ProductResponse{
		ID:          int(product.ID),
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
	})
	if err != nil {
		e := errors.FromError(err)
		return dto.ProductResponse{}, echo.NewHTTPError(int(e.GetCode()), e.GetDetail())
	}
	return dto.ProductResponse{
		ID:          int(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       int(product.Price),
		CreatedAt:   dto.DateTime(product.CreatedAt.AsTime()),
		UpdatedAt:   dto.DateTime(product.UpdatedAt.AsTime()),
	}, nil
}

func (service *productServiceImpl) Delete(ctx context.Context, productID int) error {
	_, err := service.ProductRPC.Delete(ctx, &pb.ProductID{ID: int64(productID)})
	if err != nil {
		e := errors.FromError(err)
		return echo.NewHTTPError(int(e.GetCode()), e.GetDetail())
	}
	return nil
}
