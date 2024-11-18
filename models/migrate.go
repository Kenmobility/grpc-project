package models

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Order{})
	db.AutoMigrate(&User{})
}
