package models

import (
	"encoding/json"
	"fmt"
	"gin-vue-microBlog/conf"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func readConfig() *conf.Config {
	// 使用ymal文件做读取操作
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./conf/config.yaml")
	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	var _config *conf.Config
	//将配置文件读到结构体中
	err = viper.Unmarshal(&_config)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(_config.DataBase)
	return _config
}

func InitDb() {
	// 读取ymal配置文件
	configInfo := readConfig()
	dataBaseInfo := configInfo.DataBase
	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%dms",
		dataBaseInfo.UserName,
		dataBaseInfo.PassWord,
		dataBaseInfo.Host,
		dataBaseInfo.Port,
		dataBaseInfo.User,
		dataBaseInfo.TimeOut,
	)
	var err error
	// DB的赋值要注意是全局变量还是局部变量
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("errrrr")
	}
	// defer DB.Close()
	sqlDb, _ := DB.DB()
	// 关闭数据库链接
	// defer sqlDb.Close()
	sqlDb.SetMaxIdleConns(dataBaseInfo.MaxConn) //设置最大连接数
	sqlDb.SetMaxOpenConns(dataBaseInfo.MaxOpen) //设置最大的空闲连接数
	sqlDb.Stats()
	data, err := json.Marshal(sqlDb.Stats()) //获得当前的SQL配置情况
	// 打印配置情况
	fmt.Println(string(data))
}