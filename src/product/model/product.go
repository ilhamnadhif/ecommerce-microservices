package model

import "time"

type Product struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Product) TableName() string {
	return "products"
}
