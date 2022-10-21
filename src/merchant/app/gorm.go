package app

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"merchant/config"
	"merchant/helper"
	"merchant/model"
)

func InitGorm() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config.Database.Username, config.Config.Database.Password, config.Config.Database.HostPort, config.Config.Database.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.LogFatalIfError(err)

	db.AutoMigrate(&model.Merchant{})
	return db
}
