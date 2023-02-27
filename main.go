package main

// import本项目的其他文件的包的时候，需要以项目名/包名的方式引用，并且go.mod文件的最上面需要声明本项目的module 如：module gin-vue-microBlog
import (
	"gin-vue-microBlog/routers"

	"gin-vue-microBlog/models"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// 初始化数据库设置
	models.InitDb()
}

func main() {
	// 初始化路由
	r := routers.InitRouter()
	r.Run(":8081")
}
