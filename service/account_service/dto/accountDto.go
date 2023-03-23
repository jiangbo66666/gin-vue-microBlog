package dto

import "time"

// 来自前端参数结构
type Account struct {
	Id          int
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
}

type LoginInfo struct {
	PhoneNumber string `json:"phone"`
	AccountName string `json:"accountName"`
	PassWord    string `json:"passWord"`
}

// 返回给前端的数据
type UserDetails struct {
	AccountName string    `json:"accountName"`
	PhoneNumber string    `json:"phoneNumber"`
	RecentLogin time.Time `json:"recentLogin"`
}
