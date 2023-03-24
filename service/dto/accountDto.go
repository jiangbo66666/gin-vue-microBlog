package dto

import (
	"time"
)

type Register struct {
	AccountName string `json:"accountName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

// 来自前端参数结构
type Account struct {
	Id          int
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
	AccountName string `json:"accountName"`
	Password    string `json:"Password"`
}

// 返回给前端的数据
type UserDetails struct {
	AccountName string    `json:"accountName"`
	PhoneNumber string    `json:"phoneNumber"`
	RecentLogin time.Time `json:"recentLogin"`
}
