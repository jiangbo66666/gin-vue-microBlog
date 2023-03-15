package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	AccountName string `json:"accountName"`
	jwt.RegisteredClaims
}

// 自定义秘钥
var mySecret = []byte("jiangbo")

// 生成token
func GenerateToken(Name string) (string, error) {
	claims := MyClaims{
		Name,
		jwt.RegisteredClaims{
			// 过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			//签名
			Issuer: "microBlog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 自定义秘钥
	tokenStr, err := token.SignedString(mySecret)
	if err == nil {
		return tokenStr, err
	}
	return "", errors.New("出错了")
}

func VarifyToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 校验秘钥
		return mySecret, nil
	})
	// 校验失败的时候即返回失败
	if err != nil {
		return "", err
	}
	// 解析账号名称
	claims, ok := token.Claims.(*MyClaims)
	if ok {
		return claims.AccountName, nil
	} else {
		return "", errors.New("解析账号名称失败")
	}
}
