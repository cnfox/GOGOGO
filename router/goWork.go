package router

import (
	"GOGOGO/controller/goWork"
	"github.com/gin-gonic/gin"
)

// go面试题
func goWorkRoute(router *gin.Engine) {
	workApi := router.Group("api/work")
	{
		workApi.GET("daydayup", goWork.IndexFunc) //Hello World
	}
}
