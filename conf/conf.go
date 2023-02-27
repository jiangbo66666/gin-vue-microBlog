package conf

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// host: 127.0.0.1:3306
// port: 3306
// userName: root
// passWord: imm62611
// timeOut: 1000
// 读取ymal要映射到最外层，最外层也要在项目下面
type Config struct {
	DataBase DataBase `ymal:"dataBase"`
}

type DataBase struct {
	Host     string `ymal:"host"`
	Port     string `ymal:"port"`
	UserName string `ymal:"userName"`
	PassWord string `ymal:"passWord"`
	TimeOut  int    `ymal:"timeOut"`
	MaxConn  int
	MaxOpen  int
}

func readConfig() *Config {
	// 使用ymal文件做读取操作
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./conf/config.yaml")
	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	var _config *Config
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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/gorm?charset=utf8mb4&parseTime=True&loc=Local&timeout=%dms",
		dataBaseInfo.UserName,
		dataBaseInfo.PassWord,
		dataBaseInfo.Host,
		dataBaseInfo.Port,
		dataBaseInfo.TimeOut,
	)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("errrrr")
	}
	sqlDb, _ := DB.DB()
	// 关闭数据库链接
	defer sqlDb.Close()
	sqlDb.SetMaxIdleConns(dataBaseInfo.MaxConn) //设置最大连接数
	sqlDb.SetMaxOpenConns(dataBaseInfo.MaxOpen) //设置最大的空闲连接数
	// sqlDb.Stats()
	data, err := json.Marshal(sqlDb.Stats()) //获得当前的SQL配置情况
	// 打印配置情况
	fmt.Println(string(data))
}
