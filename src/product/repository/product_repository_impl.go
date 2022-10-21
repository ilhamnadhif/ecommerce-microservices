package repository

import (
	"context"
	"gorm.io/gorm"
	"product/model"
)

func NewProductRepository() ProductRepository {
	return &productRepositoryImpl{}
}

type productRepositoryImpl struct{} // store in database

func (repository *productRepositoryImpl) FindOneByID(ctx context.Context, db *gorm.DB, productID int) (model.Product, error) {
	var product model.Product
	err := db.WithContext(ctx).First(&product, "id = ?", productID).Error
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (repository *productRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]model.Product, error) {
	var products []model.Product
	err := db.WithContext(ctx).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, err
}

func (repository *productRepositoryImpl) FindAllByMerchantID(ctx context.Context, db *gorm.DB, merchantID int) ([]model.Product, error) {
	var products []model.Product
	err := db.WithContext(ctx).Find(&products, "merchant_id = ?", merchantID).Error
	if err != nil {
		return nil, err
	}
	return products, err
}

func (repository *productRepositoryImpl) Create(ctx context.Context, db *gorm.DB, request model.Product) (model.Product, error) {
	err := db.WithContext(ctx).Create(&request).Error
	if err != nil {
		return model.Product{}, err
	}
	return request, nil
}

func (repository *productRepositoryImpl) Update(ctx context.Context, db *gorm.DB, request model.Product) (model.Product, error) {
	err := db.WithContext(ctx).Where(&model.Product{ID: request.ID}).Updates(&request).Error
	if err != nil {
		return model.Product{}, err
	}
	return request, nil
}

func (repository *productRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, productID int) error {
	err := db.WithContext(ctx).Where(&model.Product{ID: productID}).Delete(&model.Product{}).Error
	if err != nil {
		return err
	}
	return nil
}
