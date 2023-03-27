package models

import (
	"errors"

	"gorm.io/gorm"
)

// 创建用户
func CreateAccount(acc *AccountInfo) (uint, error) {
	user := acc.User
	// 先创建空用户信息表
	err := DB.Create(&user).Error
	if err != nil {
		return 0, err
	}
	acc.UserId = user.ID
	err = DB.Create(&acc).Error
	if err != nil {
		return 0, err
	}
	return acc.ID, nil
}

// 根据用户id查询账号信息
func AccountInfoById(id uint) (AccountInfo, error) {
	curUser := AccountInfo{}
	// sql语句取值且放入curUser结构体
	err := DB.Where("id = ?", id).Scopes(ActiveAccount).First(&curUser).Error
	if err != nil {
		return curUser, err
	}
	// scan将数据库查询出来的数据扫描到与前端交互的结构体中
	if err != nil && err != gorm.ErrRecordNotFound {
		return curUser, err
	}
	return curUser, err
}

// 根据账号名查询账号信息
func AccountInfoByName(name string) (AccountInfo, error) {
	curUser := AccountInfo{}
	// sql语句取值且放入curUser结构体 ,并且筛选掉停用的账号
	err := DB.Model(&AccountInfo{}).Scopes(ActiveAccount).Where("account_name = ?", name).Scan(&curUser).Error
	// scan将数据库查询出来的数据扫描到与前端交互的结构体中
	if err != nil && err != gorm.ErrRecordNotFound {
		return curUser, err
	}
	return curUser, err
}

// 根据手机号查询用户信息
func AccountInfoByPhone(phoneNumber string) (AccountInfo, error) {
	curUser := AccountInfo{}
	// sql语句取值且放入curUser结构体,并且筛选掉停用的账号
	// find方法和scan方法一样，是把数据扫描直接扫描到对应的数据中的
	err := DB.Model(&AccountInfo{}).Debug().Where("phone_number = ?", phoneNumber).Scopes(ActiveAccount).First(&curUser).Error
	// scan将数据库查询出来的数据扫描到与前端交互的结构体中
	// gorm.ErrRecordNotFound表示没有接收到数据
	if err != nil {
		return curUser, err
	}
	return curUser, err
}

// 更新账号信息
func UpdateAccountInfo(acc *AccountInfo, a map[string]interface{}) error {
	// Scopes方法可入参ScopeFunc函数用于指定统一内容
	err := DB.Model(acc).Scopes(ActiveAccount).Updates(a).Error
	if err != nil {
		return errors.New("更新失败")
	}
	return nil
}
