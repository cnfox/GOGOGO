package goCode

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

/*
	var ch = make(chan int,5)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ch chan int ,wg *sync.WaitGroup) {
		for i:=1 ;i < 5 ;i++ {
			ch <- i
		}

		fmt.Println("push over")
		defer wg.Done()
	}(ch,&wg)
	wg.Wait()

	for  {
		select {
		case i ,ok := <-ch:
			if ok {
				fmt.Println(i)
				fmt.Println("-")
			}else {
				fmt.Println("channel over")
				os.Exit(0)
			}

		case <- time.After(time.Second):
			fmt.Println("timeout")
			os.Exit(1)
		}
	}
*/

func ChannelDemo1(c *gin.Context) {
	runtime.GOMAXPROCS(1)

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		fmt.Println("A")
		wg.Done()
	}()
	go func() {
		fmt.Println("B")
		wg.Done()
	}()
	go func() {
		fmt.Println("C")
		wg.Done()
	}()
	wg.Wait()
	//1. runtime.GOMAXPROCS()
	//GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns the previous setting.
	// 设置同时可执行的cpu数量，当设置为1时，可以认为是单线程
	//2. sync.WaitGroup{}
	//一种等待同步机制，wg.Wait()会一直等待wg的值为0才会继续执行。wg.Add()增加，wg.Done()减一。
}

func ChannelDemo2(c *gin.Context) {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball)
	time.Sleep(10 * time.Second)
	<-table
	close(table)
}

type Ball struct {
	hits int
}

func player(name string, table chan *Ball) {
	for {
		ball, ok := <-table
		if !ok {
			break
		}
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(1 * time.Second)
		table <- ball
	}
}

//trace 记录了运行时的信息，能提供可视化的 Web 页面。
func TraceDemo(c *gin.Context) {
	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//main
	fmt.Println("Hello World")
}

type sub struct {
	closing chan chan error
	updates chan string
}

func (s *sub) Close() error {
	errc := make(chan error)
	s.closing <- errc
	return <-errc
}

func (s *sub) loop() {
	var err error

	for {
		select {
		case errc := <-s.closing:
			errc <- err
			close(s.updates)
			return
		}
	}
}
