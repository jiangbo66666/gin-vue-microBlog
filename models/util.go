package models

import "gorm.io/gorm"

func ActiveAccount(db *gorm.DB) *gorm.DB {
	return db.Where("status = ?", 1)
}
