package api

import (
	"encoding/json"
	"fmt"
	"gin-vue-microBlog/service"
	"gin-vue-microBlog/service/dto"
	"gin-vue-microBlog/util"

	"github.com/gin-gonic/gin"
)

// 注册账号路由函数
func RegisterAccount(ctx *gin.Context) {
	var registerInfo dto.Register
	err := bindJson(ctx, &registerInfo)
	if err != nil {
		ctx.JSON(200, Response{
			Msg:  err.Error(),
			Code: 500,
		})
		return
		// 校验账号名和密码是否都输入了
	} else if registerInfo.AccountName != "" && registerInfo.Password != "" {
		registerInfo.Password, err = util.PasswordHash(registerInfo.Password)
		if err != nil {
			ctx.JSON(200, Response{
				Msg:  err.Error(),
				Code: 500,
			})
			return
		}
		id, err := service.RegisterByAccountName(&registerInfo)
		if err != nil {
			ctx.JSON(200, Response{
				Msg:  err.Error(),
				Code: 500,
			})
			return
		}
		ctx.JSON(200, Response{
			Msg:  "注册成功",
			Code: 200,
			Data: id,
		})
	}
}

func AccountDetail(ctx *gin.Context) {
	// 简单的路由，承接数据并且发送出去
	var userInfo dto.Account
	userInfo.AccountName = ctx.MustGet("AccountName").(string)
	data, err := service.GetUserDetails(&userInfo)

	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
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

// 账号名登录路由函数
func LoginByName(c *gin.Context) {
	var loginInfo dto.Account
	bindJson(c, &loginInfo)
	token, err := service.LoginByNameAndToken(&loginInfo)
	if err == nil {
		res := Response{
			Msg:  "ok",
			Code: 200,
			Data: gin.H{"token": token},
		}
		c.JSON(200, res)
	} else {
		res := Response{
			Msg:  err.Error(),
			Code: 500,
		}
		c.JSON(200, res)
	}
}

// 绑定手机号码
func BindPhoneNumber(c *gin.Context) {
	var accPhone dto.AccountPhone
	AccountName := c.MustGet("AccountName").(string)
	bindJson(c, &accPhone)
	err := service.BindPhone(AccountName, accPhone.PhoneNumber)
	if err != nil {
		res := Response{
			Msg:  err.Error(),
			Code: 500,
		}
		c.JSON(200, res)
		return
	}
	res := Response{
		Msg:  "ok",
		Code: 200,
	}
	c.JSON(200, res)
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
