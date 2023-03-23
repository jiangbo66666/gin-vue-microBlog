package models

import (
	"gorm.io/gorm"
)

// 负责映射到前端的数据结构，需要与数据库数据分离
// 属性名保持与数据库表创建的结构体相同

func AccountInfoById(id int) (AccountInfo, error) {
	curUser := AccountInfo{}
	// sql语句取值且放入curUser结构体
	err := DB.Where("id = ?", id).Where("status = ?", 0).First(&AccountInfo{}).Error
	if err != nil {
		return curUser, err
	}
	// scan将数据库查询出来的数据扫描到与前端交互的结构体中
	err = DB.Model(&AccountInfo{}).Scan(&curUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return curUser, err
	}
	return curUser, err
}

func AccountInfoByName(name string) (AccountInfo, error) {
	curUser := AccountInfo{}
	// sql语句取值且放入curUser结构体 ,并且筛选掉停用的账号
	err := DB.Debug().Model(&AccountInfo{}).Where("account_name = ?", name).Where("status = ?", 0).Scan(&curUser).Error
	// scan将数据库查询出来的数据扫描到与前端交互的结构体中
	if err != nil && err != gorm.ErrRecordNotFound {
		return curUser, err
	}
	return curUser, err
}

func AccountInfoByPhone(phoneNumber string) (AccountInfo, error) {
	curUser := AccountInfo{}
	// sql语句取值且放入curUser结构体,并且筛选掉停用的账号
	err := DB.Model(&AccountInfo{}).Where("phone_number = ?", phoneNumber).Where("status = ?", 0).Scan(&curUser).Error
	// scan将数据库查询出来的数据扫描到与前端交互的结构体中
	if err != nil && err != gorm.ErrRecordNotFound {
		return curUser, err
	}
	return curUser, err
}
