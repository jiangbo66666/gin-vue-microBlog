package service

import (
	"errors"
	"gin-vue-microBlog/models"
	"gin-vue-microBlog/service/dto"
	"gin-vue-microBlog/util"
)

// 根据id查找账号信息
func GetAccountInfoById(user *dto.Account) (interface{}, error) {
	// 数据处理层，拿到sql的数据进行处理

	return models.AccountInfoById(user.Id)
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
	if util.PasswordVerify(login.PassWord, userInfo.Password) {
		return token, nil
	} else {
		return "", errors.New("密码错误")
	}

}
