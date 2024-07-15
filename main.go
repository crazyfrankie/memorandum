package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"memorandum/repository/cache"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"memorandum/config"
	"memorandum/pkg/util"
	"memorandum/repository/db/dao"
	"memorandum/routes"
)

func main() {
	// 加载配置并初始化
	loading()

	router := routes.NewRouter()

	// 创建一个httpserver实例
	srv := &http.Server{
		Addr:    config.HttpPort,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	fmt.Println("Server is running on http://localhost:9090")

	// 创建通道监听信号
	quit := make(chan os.Signal, 1)

	// 监听信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 阻塞直到收到信号
	<-quit
	fmt.Println("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	// 优雅地关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	fmt.Println("Server exited gracefully")
}

func loading() {
	// 解析配置
	config.Init()

	// 初始化数据库
	dao.InitDB()

	// 初始化redis
	cache.InitRedis()

	// 初始化日志
	util.InitLog()
}
