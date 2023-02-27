package models

type User struct {
	Id   int
	Name string `gorm:"name"`
	Age  int    `gorm:"age"`
}

func UserInfoSql() User {
	curUser := User{}
	// sql语句取值且放入curUser结构体
	DB.Raw("SELECT * FROM users WHERE id = ?", 1).Scan(&curUser)
	return curUser
}
