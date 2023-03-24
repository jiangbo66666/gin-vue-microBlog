package models

import "time"

type Gender string

const (
	Male   Gender = "男"
	Female Gender = "女"
	other  Gender = "其他"
)

// 用户信息表
// gorm:"-"表示该字段不会被GORM映射到数据库中
type UserInfo struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;comment:'用户id';uniqueIndex"` //id主键，自增
	Name        string    `gorm:"comment:'用户姓名'"`
	Sex         Gender    `gorm:"default:'其他';type:enum('男', '女', '其他');comment:'用户性别'"` //创建枚举类型
	BirthDay    time.Time `gorm:"default:null;comment:'用户生日'"`
	PhoneNumber string    `gorm:"comment:'用户手机号码'"`
	Email       string    `gorm:"comment:'用户邮箱'"`
	Address     string    `gorm:"comment:'用户地址'"`
	CreateBy    int       `gorm:"comment:'用户由谁创建'"`
	CreateAt    time.Time `gorm:"default:(NOW());comment:'创建账号时间'"` //创建默认时间
	UpdateAt    time.Time `gorm:"default:(NOW());comment:'更新时间'"`
	RecentLogin time.Time `gorm:"default:(NOW());comment:'最近登陆时间'"`
	HeaderImage string    `gorm:"comment:'头像地址'"`
	Profile     string    `gorm:"comment:'个人简介'"`
}

// 唯一索引需要设置为not null null在数据库中是不被认为重复的。
type AccountInfo struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;comment:'账号id'"` //id主键，自增
	UserId      uint      `gorm:"comment:'用户id'"`
	PhoneNumber string    `gorm:"not null;comment:'用户手机号码';type:varchar(20)"` //账号密码表带入手机号码，方便登录流程，减少登录查询
	AccountName string    `gorm:"not null;comment:'账号名';uniqueIndex;type:varchar(20)"`
	Password    string    `gorm:"comment:'账号密码'"`
	Email       string    `gorm:"comment:'用户注册邮箱'"`
	CreateAt    time.Time `gorm:"default:(NOW());comment:'创建账号时间'"`
	UpdateAt    time.Time `gorm:"comment:更新时间;default:(NOW())"`
	RecentLogin time.Time `gorm:"comment:最近登录时间;default:(NOW())"`
	Status      int       `gorm:"default:0;comment:'用户账号状态'"`
	User        UserInfo  `gorm:"foreignKey:UserId"` //外键关联userInfo表
}
