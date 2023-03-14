package account_service

import "gin-vue-microBlog/models"

// 前端参数结构
type Account struct {
	Id          int
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
}

// 根据id查找账号信息
func (user *Account) GetAccountInfoById() (models.AccountInfoJson, error) {
	// 数据处理层，拿到sql的数据进行处理
	return models.AccountInfoById(user.Id)
}

// 根据账号名称查找用户信息
func (user *Account) GetAccountInfoByName() (models.AccountInfoJson, error) {
	// 数据处理层，拿到sql的数据进行处理
	return models.AccountInfoByName(user.Name)
}

// 根据手机号码查找账号信息
func (user *Account) GetAccountInfoByPhone() (models.AccountInfoJson, error) {
	// 数据处理层，拿到sql的数据进行处理
	return models.AccountInfoByPhone(user.PhoneNumber)
}
