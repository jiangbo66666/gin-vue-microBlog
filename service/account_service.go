package service

import (
	"errors"
	"gin-vue-microBlog/models"
	"gin-vue-microBlog/service/dto"
	"gin-vue-microBlog/util"
)

// 注册账号
func RegisterByAccountName(registerInfo *dto.Register) (uint, error) {
	// 根据账号名称注册账号
	accountInfo := models.AccountInfo{
		AccountName: registerInfo.AccountName,
		Password:    registerInfo.Password,
	}
	// 先查询账号是否存在于数据库，若存在则不创建
	info, err := models.AccountInfoByName(accountInfo.AccountName)
	if err != nil {
		return 0, errors.New("查询数据库失败")
	}
	if info.AccountName == accountInfo.AccountName {
		return 0, errors.New("账号已存在")
	} else {
		id, err := models.CreateAccount(&accountInfo)
		return id, err
	}
}

// 根据id查找账号信息
func GetAccountInfoById(id uint) (interface{}, error) {
	// 数据处理层，拿到sql的数据进行处理

	return models.AccountInfoById(id)
}

// 根据账号名称查找用户信息
func GetUserDetails(user *dto.Account) (interface{}, error) {
	// 数据处理层，拿到sql的数据进行处理
	userDetail, err := models.AccountInfoByName(user.AccountName)
	data := dto.UserDetails{
		AccountName: userDetail.AccountName,
		PhoneNumber: userDetail.PhoneNumber,
		RecentLogin: userDetail.RecentLogin,
	}
	return data, err
}

// 根据手机号码查找账号信息
func GetAccountInfoByPhone(user *dto.Account) (interface{}, error) {
	// 数据处理层，拿到sql的数据进行处理
	return models.AccountInfoByPhone(user.PhoneNumber)
}

func LoginByNameAndToken(login *dto.Account) (string, error) {
	userInfo, err := models.AccountInfoByName(login.AccountName)
	if err != nil {
		return "", err
	}
	// 根据账号名称生成token
	token, err := util.GenerateToken(userInfo.AccountName)
	if err != nil {
		return "", err
	}
	if util.PasswordVerify(login.Password, userInfo.Password) {
		return token, nil
	} else {
		return "", errors.New("密码错误")
	}

}
