package model

import "time"

type Merchant struct {
	ID        int    `gorm:"primaryKey"`
	Email     string `gorm:"type:varchar(50);unique"`
	Password  string `gorm:"type:varchar(50)"`
	Name      string `gorm:"type:varchar(50);unique"`
	Slug      string `gorm:"type:varchar(50)"`
	Balance   int
	ImageID   int
	Image     string `gorm:"type:varchar(150)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Merchant) TableName() string {
	return "merchants"
}
