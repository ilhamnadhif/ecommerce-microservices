package repository

import (
	"context"
	"gorm.io/gorm"
	"product/model"
)

type ProductRepository interface {
	FindOneByID(ctx context.Context, db *gorm.DB, productID int) (model.Product, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]model.Product, error)
	FindAllByMerchantID(ctx context.Context, db *gorm.DB, merchantID int) ([]model.Product, error)
	Create(ctx context.Context, db *gorm.DB, request model.Product) (model.Product, error)
	Update(ctx context.Context, db *gorm.DB, request model.Product) (model.Product, error)
	Delete(ctx context.Context, db *gorm.DB, productID int) error
}
