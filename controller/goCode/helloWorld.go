package goCode

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	fmt.Println("Hello World !")
	fmt.Println("你好 世界 !")
}
