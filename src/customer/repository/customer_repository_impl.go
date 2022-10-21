package repository

import (
	"context"
	"customer/model"

	"gorm.io/gorm"
)

func NewCustomerRepository() CustomerRepository {
	return &customerRepositoryImpl{}
}

type customerRepositoryImpl struct{} // store in database

func (repository *customerRepositoryImpl) FindOneByID(ctx context.Context, db *gorm.DB, customerID int) (model.Customer, error) {
	var customer model.Customer
	err := db.WithContext(ctx).First(&customer, "id = ?", customerID).Error
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func (repository *customerRepositoryImpl) FindOneByEmail(ctx context.Context, db *gorm.DB, email string) (model.Customer, error) {
	var customer model.Customer
	err := db.WithContext(ctx).First(&customer, "email = ?", email).Error
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func (repository *customerRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]model.Customer, error) {
	var customers []model.Customer
	err := db.WithContext(ctx).Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, err
}

func (repository *customerRepositoryImpl) Create(ctx context.Context, db *gorm.DB, request model.Customer) (model.Customer, error) {
	err := db.WithContext(ctx).Create(&request).Error
	if err != nil {
		return model.Customer{}, err
	}
	return request, nil
}

func (repository *customerRepositoryImpl) Update(ctx context.Context, db *gorm.DB, request model.Customer) (model.Customer, error) {
	err := db.WithContext(ctx).Where(&model.Customer{ID: request.ID}).Updates(&request).Error
	if err != nil {
		return model.Customer{}, err
	}
	return request, nil
}

func (repository *customerRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, customerID int) error {
	err := db.WithContext(ctx).Where(&model.Customer{ID: customerID}).Delete(&model.Customer{}).Error
	if err != nil {
		return err
	}
	return nil
}
