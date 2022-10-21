package model

import "time"

type Merchant struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Merchant) TableName() string {
	return "merchants"
}
