package user_service

import "gin-vue-microBlog/models"

// 前端参数结构
type User struct {
	Id   int
	Name string
	Age  int
}

func (user *User) GetUserInfoById() (models.AccountInfoJson, error) {
	// 数据处理层，拿到sql的数据进行处理
	return models.UserInfoById(user.Id)
}
