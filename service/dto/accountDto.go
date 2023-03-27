package dto

import (
	"gin-vue-microBlog/models"
	"time"
)

// 来自前端的注册信息
type Register struct {
	AccountName string `json:"accountName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

// 登录的时候传递的信息
type Account struct {
	Id          int
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
	AccountName string `json:"accountName"`
	Password    string `json:"Password"`
}

// 查询用户详情返回的信息
type AccountDetails struct {
	AccountName string    `json:"accountName"`
	PhoneNumber string    `json:"phoneNumber"`
	RecentLogin time.Time `json:"recentLogin"`
	UserDetails models.UserInfo
}

type UserDetails struct {
}

type AccountPhone struct {
	PhoneNumber string `json:"phoneNumber"`
}
