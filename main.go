package main

import (
	"context"
	"file-manager/global"
	"file-manager/initialize"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	initialize.InitConfig("./conf/config.yaml") // 初始化Viper
	router := initialize.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.CONFIG.Application.Port),
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf(`欢迎使用 SiriX开发框架,程序运行地址:http://127.0.0.1:%d`, global.CONFIG.Application.Port)
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			os.Exit(0)
		}
	}()

	shutDown(s)
}

func shutDown(s *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutdown Server ...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		//global.LOG.Fatal("Server Shutdown:" + err.Error())
	}
	fmt.Println("Server exiting")
}
