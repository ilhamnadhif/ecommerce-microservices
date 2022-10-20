package model

import "time"

type Customer struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Customer) TableName() string {
	return "customers"
}
