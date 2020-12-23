package main

import (
	"GOGOGO/controller/common"
	"GOGOGO/router"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	//"github.com/DeanThompson/ginpprof"

	"GOGOGO/config"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	fmt.Println("gin.Version: ", gin.Version)

	// Creates a router without any middleware by default
	app := gin.New()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	maxSize := int64(config.ServerConfig.MaxMultipartMemory)
	app.MaxMultipartMemory = maxSize << 40 // 3 MiB

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	app.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	app.Use(gin.Recovery())
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

	// 旧的请求等待30s后关闭，拒绝新的请求
	// TODO: Supervisor的stopwaitsecs参数默认为10s，表示如果进程超过10s还未正常关闭，则发送SIGKILL强制关闭
	ctx, closeServer := context.WithTimeout(context.Background(), 30*time.Second)
	defer closeServer()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("Server Shutdown:", err)
	}
	log.Println("Server gracefully stopped")

	//app.Run(":" + fmt.Sprintf("%d", config.ServerConfig.Port))
}
