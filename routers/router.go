package routers

import (
	"gin-vue-microBlog/routers/api"
	"gin-vue-microBlog/util"

	"github.com/gin-gonic/gin"
)

// 返回gin的路由类型
func InitRouter() *gin.Engine {
	// 初始化路由
	r := gin.New()
	// log日志
	r.Use(gin.Logger())
	//
	r.Use(gin.Recovery())

	// 静态文件读取
	// r.StaticFS()
	// 中间件，路由守卫
	r.Use(func(ctx *gin.Context) {

	})
	r.POST("/register", api.RegisterAccount)

	r.POST("/login", api.LoginByName)

	// 路由分组,读取userdetail handler
	user := r.Group("/api/user")
	// 使用jwt
	// user.Use(jwt.JWT())
	{
		user.Use(func(ctx *gin.Context) {
			token := ctx.GetHeader("Token")
			//校验账号信息
			AccountName, err := util.VarifyToken(token)
			res := api.Response{
				Msg:  "登录失效",
				Code: 501,
			}
			if err != nil {
				ctx.JSON(200, res)

				ctx.Abort()
				return
			} else {
				// token校验通过，将token中的账号信息存储起来
				ctx.Set("AccountName", AccountName)
				ctx.Next()
			}
		})
		user.GET("/info", api.UserDetail)
	}

	return r
}
