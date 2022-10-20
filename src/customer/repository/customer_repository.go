package repository

import (
	"context"
	"customer/model"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindOneByID(ctx context.Context, db *gorm.DB, userID int) (model.Customer, error)
	FindOneByEmail(ctx context.Context, db *gorm.DB, email string) (model.Customer, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]model.Customer, error)
	Create(ctx context.Context, db *gorm.DB, request model.Customer) (model.Customer, error)
	Update(ctx context.Context, db *gorm.DB, request model.Customer) (model.Customer, error)
	Delete(ctx context.Context, db *gorm.DB, userID int) error
}
