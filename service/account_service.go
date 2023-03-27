package service

import (
	"errors"
	"gin-vue-microBlog/models"
	"gin-vue-microBlog/service/dto"
	"gin-vue-microBlog/util"
	"time"

	"gorm.io/gorm"
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

// 根据账号名登录且返回token字符串
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
		userInfo.RecentLogin = time.Now()
		// 根据账号更新指定的内容
		_ = models.UpdateAccountInfo(&userInfo, map[string]interface{}{
			"RecentLogin": time.Now(),
		})
		return token, nil
	} else {
		return "", errors.New("密码错误")
	}
}

// 绑定手机号码
func BindPhone(accountName string, phoneNumber string) error {
	NameuserInfo, err := models.AccountInfoByName(accountName)
	if err != nil {
		return errors.New("账号不存在")
	}
	if NameuserInfo.PhoneNumber != "" {
		return errors.New("该账号已绑定手机号")
	}
	_, err = models.AccountInfoByPhone(phoneNumber)
	// 若没有查询到，则可以实施绑定操作
	if err == gorm.ErrRecordNotFound {
		err = models.UpdateAccountInfo(&NameuserInfo, map[string]interface{}{
			"PhoneNumber": phoneNumber,
		})
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("该手机号已被其他账号绑定")
	}
}
