package main

// import本项目的其他文件的包的时候，需要以项目名/包名的方式引用，并且go.mod文件的最上面需要声明本项目的module 如：module gin-vue-microBlog
import (
	"encoding/json"
	"fmt"
	"gin-vue-microBlog/conf"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
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
	fmt.Println(dataBaseInfo.MaxConn)
	fmt.Println(dataBaseInfo.MaxOpen)
	sqlDb.SetMaxIdleConns(dataBaseInfo.MaxConn) //设置最大连接数
	sqlDb.SetMaxOpenConns(dataBaseInfo.MaxOpen) //设置最大的空闲连接数
	// sqlDb.Stats()
	data, err := json.Marshal(sqlDb.Stats()) //获得当前的SQL配置情况
	// 打印配置情况
	fmt.Println(string(data))

}

func main() {
	router := gin.Default()

	router.Run(":80")
}

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
