package user_service

import "gin-vue-microBlog/models"

type User struct {
	Id   int
	Name string `gorm:"name"`
	Age  int    `gorm:"age"`
}

func (user *User) GetUserInfoById() (models.User, error) {
	// 数据处理层，拿到sql的数据进行处理
	return models.UserInfoById(user.Id)
}
