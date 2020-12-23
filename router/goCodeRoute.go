package router

import (
	"GOGOGO/controller/goCode"
	"github.com/gin-gonic/gin"
)

// go基本语法
func goCodeRoute(router *gin.Engine) {
	codeApi := router.Group("api/code")
	{
		codeApi.GET("helloWorld", goCode.HelloWorld) //Hello World
		codeApi.GET("dataType", goCode.DataType)     //数据类型

		codeApi.GET("goroutineDemo", goCode.GoroutineDemo) //goroutine
		codeApi.GET("gchannelDemo1", goCode.ChannelDemo1)  //channel1
		codeApi.GET("gchannelDemo2", goCode.ChannelDemo2)  //channel2

		codeApi.GET("trace", goCode.TraceDemo) //trace

	}
}
