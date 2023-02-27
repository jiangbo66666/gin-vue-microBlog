package api

import (
	"gin-vue-microBlog/service/user_service"

	"github.com/gin-gonic/gin"
)

func UserDetail(ctx *gin.Context) {
	// 简单的路由，承接数据并且发送出去
	ctx.JSON(200, gin.H{
		"msg":  "okok",
		"data": user_service.GetUserInfo(),
	})
}
