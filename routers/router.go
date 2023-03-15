/*
 * @Author: jiangbo jiangbo1996@outlook.com
 * @Date: 2023-02-27 14:11:47
 * @LastEditors: jiangbo jiangbo1996@outlook.com
 * @LastEditTime: 2023-02-27 15:02:42
 * @FilePath: \gin-vue-microBlog\routers\router.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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
			if err != nil {
				ctx.JSON(200, gin.H{
					"code": 501,
					"msg":  "登录失效",
				})

				ctx.Abort()
				return
			} else {
				// 将token中的账号信息存储起来
				ctx.Set("AccountName", AccountName)
				ctx.Next()
			}
		})
		user.POST("/info", api.UserDetail)
	}

	return r
}
