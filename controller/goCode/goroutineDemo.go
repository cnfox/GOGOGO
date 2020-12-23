package goCode

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
	"sync"
)

func GoroutineDemo(c *gin.Context) {
	runtime.GOMAXPROCS(1) //设置核心数
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go PrintlnHello(&wg, i)
	}

	//go func(s string) {
	//	for i := 0; i < 2; i++ {
	//		fmt.Println(s)
	//	}
	//}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
	wg.Wait() // 等待所有登记的goroutine都结束
	fmt.Println("over")

}

func PrintlnHello(wg *sync.WaitGroup, i int) {
	wg.Done() // goroutine结束就登记-1
	fmt.Println(i)
}
