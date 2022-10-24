package model

import "time"

type Cart struct {
	ID         int `gorm:"primaryKey"`
	CustomerID int
	ProductID  int
	Quantity   int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (Cart) TableName() string {
	return "carts"
}
