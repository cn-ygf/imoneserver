package main

import (
	_ "github.com/cn-ygf/imoneserver/api"
	_ "github.com/cn-ygf/imoneserver/gateway"
	"github.com/cn-ygf/imoneserver/lib/config"
	"github.com/cn-ygf/imoneserver/service"
	_ "github.com/cn-ygf/imoneserver/session"
	"github.com/davyxu/golog"
	"os"
	"os/signal"
	"syscall"
)

var log = golog.New("imone")

func main() {
	// 加载配置文件
	if len(os.Args) < 2 {
		log.Errorln("Configure file not find!")
		return
	}
	// 加载配置文件
	err := config.Load(os.Args[1])
	if err != nil {
		log.Errorf("Configure file is load failed!Error:%s\n", err.Error())
		return
	}
	// 启动api服务
	api := service.NewService("api", "api v1")
	api.Run(config.GetString("api_bind"), config.GetString("api_version"))

	// 启动gateway服务
	// TODO
	gate := service.NewService("gateway")
	gate.Run(config.GetString("gateway_bind"))
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
		log.Infof("imoneserver get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			api.Close()
			session.Close()
			gate.Close()
			log.Infof("imoneserver exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
