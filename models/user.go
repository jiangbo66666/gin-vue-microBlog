package models

import "gorm.io/gorm"

type User struct {
	Id   int
	Name string `gorm:"name"`
	Age  int    `gorm:"age"`
}

func UserInfoById(id int) (User, error) {
	curUser := User{}
	// sql语句取值且放入curUser结构体
	err := DB.Raw("SELECT * FROM users WHERE id = ?", id).Scan(&curUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return curUser, err
	}
	return curUser, err
}
