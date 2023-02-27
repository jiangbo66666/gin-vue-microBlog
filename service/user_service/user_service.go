package user_service

import "gin-vue-microBlog/models"

func GetUserInfo() models.User {
	// 数据处理层，拿到sql的数据进行处理
	return models.UserInfoSql()
}
