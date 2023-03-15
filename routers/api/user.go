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
	userInfo.Name = ctx.MustGet("AccountName").(string)
	data, err := userInfo.GetAccountInfoByName()

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

func LoginByName(c *gin.Context) {
	var loginInfo account_service.LoginInfo
	bindJson(c, &loginInfo)
	logined, token, err := loginInfo.CheckLoginByNameAndToken()
	if logined && (err == nil) {
		c.JSON(200, gin.H{
			"code": 200,
			"data": gin.H{"token": token},
			"msg":  "登陆成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "登陆失败",
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
