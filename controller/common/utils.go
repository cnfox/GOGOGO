package common

import (
	"GOGOGO/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// int to string
func IntToStr(intvalue int) string {
	return strconv.Itoa(intvalue)
}

// SendErrJSON 有错误发生时，发送错误JSON
func SendErrJSON(msg string, args ...interface{}) {
	if len(args) == 0 {
		panic("缺少 *gin.Context")
	}
	var c *gin.Context
	var errNo = model.ErrorCode.ERROR
	if len(args) == 1 {
		theCtx, ok := args[0].(*gin.Context)
		if !ok {
			panic("缺少 *gin.Context")
		}
		c = theCtx
	} else if len(args) == 2 {
		theErrNo, ok := args[0].(int)
		if !ok {
			panic("errNo不正确")
		}
		errNo = theErrNo
		theCtx, ok := args[1].(*gin.Context)
		if !ok {
			panic("缺少 *gin.Context")
		}
		c = theCtx
	}
	fmt.Println(msg) // todo for debug, added by leek, 应该放到业务里决定是否打印错误日志 还是统一放到这里?
	c.JSON(http.StatusOK, gin.H{
		"errNo": errNo,
		"msg":   msg,
		"data":  gin.H{},
	})
	// 终止请求链
	c.Abort()
}
