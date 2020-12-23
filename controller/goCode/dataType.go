package goCode

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func DataType(c *gin.Context) {
	//强类型 健壮性
	//变量
	var z int
	z = 1
	fmt.Println(z) //必须使用

	//变量推导

	var (
		j = 0
		k = 1
	)
	fmt.Println(j, k)

	//零值

	// :=
	//指针
	//赋值

	//常量
	const name = "为什么布丁"
	const (
		one   = iota + 1
		two   = iota + 1
		three = iota + 1
		four  = iota + 1
	)

	//数据类型互相转换
	//strings
}
