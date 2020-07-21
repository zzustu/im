package initialize

import (
	"context"
	"github.com/zzustu/im/global"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
)

func Start() {

	// 加载配置文件信息, 初始化配置
	if err := loadConfig(); err != nil {
		log.Fatal(err)
	}

	// 配置日志信息
	if err := configLog(); err != nil {
		log.Fatal(err)
	}

	printInfo()

	// 初始化BoltDB
	if err := loadBoltDB(); err != nil {
		log.Fatal(err)
	}

	// 连接数据库
	if err := connectDB(); err != nil {
		log.Fatal(err)
	}

	// 启动HTTP
	srv, err := getHttpServer()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
		_ = global.DB.Close()
		_ = global.TB.Close()
	}()

	global.CM.Broadcast("系统准备关机")

	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Print("Shutdown.")
}

func printInfo() {
	log.Printf("进 程 号: %d", os.Getpid())
	log.Printf("操作系统: %s", runtime.GOOS)
	log.Printf("软件版本: %s", global.Version)
}
