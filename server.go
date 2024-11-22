package main

import (
	"context"
	"errors"
	"go-SkyTakeaway/config"
	_ "go-SkyTakeaway/middleware/timer"
	"go-SkyTakeaway/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// 初始化服务器配置信息
func init() {
	config.Init()
}

func main() {
	// 获得路由
	r := router.Router()
	// 创建服务器
	srv := &http.Server{
		Addr:    config.ServerConfig.Server.Port,
		Handler: r,
	}
	// 协程运行并监听服务器
	go func() {
		// 如果不是关闭服务器错误，打印错误并退出
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("an error occurred while running the server: %s\n", err)
		}
	}()
	// "优雅"退出系统
	// 创建缓存为1的信号量channel
	quit := make(chan os.Signal, 1)
	// 将类似ctrl+C的操作的信号量发送给信号channel
	signal.Notify(quit, os.Interrupt)
	// 阻塞等待取出信号量
	<-quit
	// 接收到信号量后打印提示
	log.Println("server is being shutdown...")
	// 创建上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 延迟调用cancel，确保上下文能取消
	defer cancel()
	// 调用Shutdown关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown error:", err)
	}
	log.Println("server has been safely shutdown")
}
