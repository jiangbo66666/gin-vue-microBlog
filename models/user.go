package models

import (
	"errors"

	"gorm.io/gorm"
)

func CreateAccount(acc *AccountInfo) (uint, error) {
	user := acc.User
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

func AccountInfoByName(name string) (AccountInfo, error) {
	curUser := AccountInfo{}
	// sql语句取值且放入curUser结构体 ,并且筛选掉停用的账号
	err := DB.Debug().Model(&AccountInfo{}).Scopes(ActiveAccount).Where("account_name = ?", name).Scan(&curUser).Error
	// scan将数据库查询出来的数据扫描到与前端交互的结构体中
	if err != nil && err != gorm.ErrRecordNotFound {
		return curUser, err
	}
	return curUser, err
}

func AccountInfoByPhone(phoneNumber string) (AccountInfo, error) {
	curUser := AccountInfo{}
	// sql语句取值且放入curUser结构体,并且筛选掉停用的账号
	err := DB.Model(&AccountInfo{}).Where("phone_number = ?", phoneNumber).Scopes(ActiveAccount).Scan(&curUser).Error
	// scan将数据库查询出来的数据扫描到与前端交互的结构体中
	if err != nil && err != gorm.ErrRecordNotFound {
		return curUser, err
	}
	return curUser, err
}

func UpdateAccountInfo(acc *AccountInfo, a map[string]interface{}) error {
	// Scopes方法可入参ScopeFunc函数用于指定统一内容
	err := DB.Model(acc).Scopes(ActiveAccount).Updates(a).Error
	if err != nil {
		return errors.New("更新失败")
	}
	return nil
}
