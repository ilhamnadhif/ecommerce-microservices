package repository

import (
	"context"
	"merchant/model"

	"gorm.io/gorm"
)

type MerchantRepository interface {
	FindOneByID(ctx context.Context, db *gorm.DB, userID int) (model.Merchant, error)
	FindOneByEmail(ctx context.Context, db *gorm.DB, email string) (model.Merchant, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]model.Merchant, error)
	Create(ctx context.Context, db *gorm.DB, request model.Merchant) (model.Merchant, error)
	Update(ctx context.Context, db *gorm.DB, request model.Merchant) (model.Merchant, error)
	Delete(ctx context.Context, db *gorm.DB, userID int) error
}
