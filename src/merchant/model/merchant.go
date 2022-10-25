package model

import "time"

type Merchant struct {
	ID        int    `gorm:"primaryKey"`
	Email     string `gorm:"type:varchar(50);not null;unique"`
	Password  string `gorm:"type:varchar(50);not null"`
	Name      string `gorm:"type:varchar(50);not null;unique"`
	Slug      string `gorm:"type:varchar(50);not null"`
	Balance   int    `gorm:"not null"`
	ImageID   int    `gorm:"not null"`
	Image     string `gorm:"type:varchar(150);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Merchant) TableName() string {
	return "merchants"
}
