package user_service

import "gin-vue-microBlog/models"

// 前端参数结构
type User struct {
	Id          int
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
}

func (user *User) GetAccountInfoById() (models.AccountInfoJson, error) {
	// 数据处理层，拿到sql的数据进行处理
	return models.AccountInfoById(user.Id)
}

func (user *User) GetAccountInfoByName() (models.AccountInfoJson, error) {
	// 数据处理层，拿到sql的数据进行处理
	return models.AccountInfoByName(user.Name)
}

func (user *User) GetAccountInfoByPhone() (models.AccountInfoJson, error) {
	// 数据处理层，拿到sql的数据进行处理
	return models.AccountInfoByPhone(user.PhoneNumber)
}
