package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserDetail(ctx *gin.Context) {
	fmt.Println("这里是用户详情")
}
