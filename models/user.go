package models

import (
	"time"

	"gorm.io/gorm"
)

// 负责映射到前端的数据结构，需要与数据库数据分离
// 属性名保持与数据库表创建的结构体相同
type AccountInfoJson struct {
	AccountName string
	PhoneNumber string
	RecentLogin time.Time
	Status      int
}

func UserInfoById(id int) (AccountInfoJson, error) {
	curUser := AccountInfoJson{}
	// sql语句取值且放入curUser结构体
	err := DB.Where("id = ?", id).Where("status = ?", 0).First(&AccountInfo{}).Error
	if err != nil {
		return curUser, err
	}
	// scan蒋数据库查询出来的数据扫描到与前端交互的结构体中
	err = DB.Model(&AccountInfo{}).Scan(&curUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return curUser, err
	}
	return curUser, err
}
