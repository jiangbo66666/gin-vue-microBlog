package models

import "gorm.io/gorm"

// 数据库操作中的共用方法整合
func ActiveAccount(db *gorm.DB) *gorm.DB {
	return db.Where("status = ?", 1)
}
