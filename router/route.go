package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	// "eaccount/controller/coop"
	"GOGOGO/controller/version"
	"strings"
)

// Route 路由
func Route(router *gin.Engine) {
	//router.Use(cors()) //使用中间件 解决跨域问题 紧开发环境使用

	//webPublic := router.Group("api/public")
	//{
	//	webPublic.POST("/login", user.Login) // for web
	//}

	//获取 当前版本 || 更新进度
	router.GET("api/getVersion", version.GetVersion)

	goCodeRoute(router)
	goFrameworkRoute(router)
	goUtilsRoute(router)
}

// 跨域问题参考文章:
// https://www.jianshu.com/p/89a377c52b48
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, Content-Type, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		var origins = []string{"http://localhost:8081", "http://localhost:8080"} // 允许跨域
		if origin != "" {
			ogn := ""
			for _, v := range origins {
				if origin == v {
					ogn = v
				}
			}

			c.Header("Access-Control-Allow-Origin", ogn)
			c.Header("Access-Control-Allow-Headers", headerStr)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, Token")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}
