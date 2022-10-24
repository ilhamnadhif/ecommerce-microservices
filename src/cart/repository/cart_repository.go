package repository

import (
	"cart/model"
	"context"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindOneByID(ctx context.Context, db *gorm.DB, cartID int) (model.Cart, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]model.Cart, error)
	FindAllByCustomerID(ctx context.Context, db *gorm.DB, customerID int) ([]model.Cart, error)
	Create(ctx context.Context, db *gorm.DB, request model.Cart) (model.Cart, error)
	Update(ctx context.Context, db *gorm.DB, request model.Cart) (model.Cart, error)
	Delete(ctx context.Context, db *gorm.DB, cartID int) error
}

func NewCartRepository() CartRepository {
	return &cartRepositoryImpl{}
}

type cartRepositoryImpl struct{} // store in database

func (repository *cartRepositoryImpl) FindOneByID(ctx context.Context, db *gorm.DB, cartID int) (model.Cart, error) {
	var cart model.Cart
	err := db.WithContext(ctx).First(&cart, "id = ?", cartID).Error
	if err != nil {
		return model.Cart{}, err
	}
	return cart, nil
}

func (repository *cartRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]model.Cart, error) {
	var carts []model.Cart
	err := db.WithContext(ctx).Find(&carts).Error
	if err != nil {
		return nil, err
	}
	return carts, err
}

func (repository *cartRepositoryImpl) FindAllByCustomerID(ctx context.Context, db *gorm.DB, customerID int) ([]model.Cart, error) {
	var carts []model.Cart
	err := db.WithContext(ctx).Find(&carts, "customer_id = ?", customerID).Error
	if err != nil {
		return nil, err
	}
	return carts, err
}

func (repository *cartRepositoryImpl) Create(ctx context.Context, db *gorm.DB, request model.Cart) (model.Cart, error) {
	err := db.WithContext(ctx).Create(&request).Error
	if err != nil {
		return model.Cart{}, err
	}
	return request, nil
}

func (repository *cartRepositoryImpl) Update(ctx context.Context, db *gorm.DB, request model.Cart) (model.Cart, error) {
	err := db.WithContext(ctx).Where(&model.Cart{ID: request.ID}).Updates(&request).Error
	if err != nil {
		return model.Cart{}, err
	}
	return request, nil
}

func (repository *cartRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, cartID int) error {
	err := db.WithContext(ctx).Where(&model.Cart{ID: cartID}).Delete(&model.Cart{}).Error
	if err != nil {
		return err
	}
	return nil
}
