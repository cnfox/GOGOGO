package main

import (
	"GOGOGO/config"
	"GOGOGO/controller/common"
	"GOGOGO/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("gin.Version: ", gin.Version)

	// Creates a router without any middleware by default
	app := gin.New()
	router.Route(app)
	fmt.Println("----------------- Prod  Prod  Prod  Prod  Prod  Prod  Prod  Prod  ----------------- ")
	srv := &http.Server{
		Addr:    ":" + common.IntToStr(config.ServerConfig.Port),
		Handler: app,
	}
	fmt.Println("appRun port:", config.ServerConfig.Port)

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待一个信号量并优雅关闭服务
	quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGUSR2) // linux
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT) //window
	<-quit
	log.Println("Shutdown Server ...")

}
