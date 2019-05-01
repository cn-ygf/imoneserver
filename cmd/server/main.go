package main

import (
	_ "github.com/cn-ygf/imoneserver/api"
	"github.com/cn-ygf/imoneserver/service"
	_ "github.com/cn-ygf/imoneserver/session"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// 启动api服务
	api := service.NewService("api", "api v1")
	api.Run("127.0.0.1:9000")

	// 启动gateway服务
	// TODO

	// 启动对象存储服务
	// TODO

	// 启动会话管理服务
	session := service.NewService("session")
	session.Run()
	// 启动IM服务
	// TODO

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("imoneserver get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			api.Close()
			session.Close()
			log.Printf("imoneserver exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
