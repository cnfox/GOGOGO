package goWork

import (
	"GOGOGO/controller/common"
	"GOGOGO/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func IndexFunc(c *gin.Context) {
	SendErrJSON := common.SendErrJSON

	var ok model.Answer
	if c.BindJSON(&ok) != nil {
		SendErrJSON("参数错误", c)
		return
	}

	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}
