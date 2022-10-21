package repository

import (
	"context"
	"merchant/model"

	"gorm.io/gorm"
)

func NewMerchantRepository() MerchantRepository {
	return &merchantRepositoryImpl{}
}

type merchantRepositoryImpl struct{} // store in database

func (repository *merchantRepositoryImpl) FindOneByID(ctx context.Context, db *gorm.DB, merchantID int) (model.Merchant, error) {
	var merchant model.Merchant
	err := db.WithContext(ctx).First(&merchant, "id = ?", merchantID).Error
	if err != nil {
		return model.Merchant{}, err
	}
	return merchant, nil
}

func (repository *merchantRepositoryImpl) FindOneByEmail(ctx context.Context, db *gorm.DB, email string) (model.Merchant, error) {
	var merchant model.Merchant
	err := db.WithContext(ctx).First(&merchant, "email = ?", email).Error
	if err != nil {
		return model.Merchant{}, err
	}
	return merchant, nil
}

func (repository *merchantRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]model.Merchant, error) {
	var merchants []model.Merchant
	err := db.WithContext(ctx).Find(&merchants).Error
	if err != nil {
		return nil, err
	}
	return merchants, err
}

func (repository *merchantRepositoryImpl) Create(ctx context.Context, db *gorm.DB, request model.Merchant) (model.Merchant, error) {
	err := db.WithContext(ctx).Create(&request).Error
	if err != nil {
		return model.Merchant{}, err
	}
	return request, nil
}

func (repository *merchantRepositoryImpl) Update(ctx context.Context, db *gorm.DB, request model.Merchant) (model.Merchant, error) {
	err := db.WithContext(ctx).Where(&model.Merchant{ID: request.ID}).Updates(&request).Error
	if err != nil {
		return model.Merchant{}, err
	}
	return request, nil
}

func (repository *merchantRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, merchantID int) error {
	err := db.WithContext(ctx).Where(&model.Merchant{ID: merchantID}).Delete(&model.Merchant{}).Error
	if err != nil {
		return err
	}
	return nil
}
