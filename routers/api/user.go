package api

import (
	"encoding/json"
	"fmt"
	"gin-vue-microBlog/service/account_service"

	"github.com/gin-gonic/gin"
)

func UserDetail(ctx *gin.Context) {
	// 简单的路由，承接数据并且发送出去
	var userInfo account_service.Account
	bindJson(ctx, &userInfo)
	data, err := userInfo.GetAccountInfoByPhone()
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "消息获取错误",
			"data": data,
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "ok",
			"data": data,
		})
	}
}

// 响应数据合并到结构体
func bindJson(c *gin.Context, struc any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err = json.Unmarshal(body, &struc)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}
