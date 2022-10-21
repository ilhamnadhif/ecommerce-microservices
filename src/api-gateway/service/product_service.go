package service

import (
	"api-gateway/dto"
	"context"
)

type ProductService interface {
	FindOneByID(ctx context.Context, productID int) (dto.ProductResponse, error)
	FindAll(ctx context.Context) ([]dto.ProductResponse, error)
	Create(ctx context.Context, request dto.ProductCreateReq) (dto.ProductResponse, error)
	Update(ctx context.Context, request dto.ProductUpdateReq) (dto.ProductResponse, error)
	Delete(ctx context.Context, productID int) error
}
