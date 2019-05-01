package session

import (
	"log"
)

// 会话管理服务
type sessionService struct {
}

func (session *sessionService) Run(param ...interface{}) {
	log.Printf("%s: service is runing\n", session.TypeName())
}

func (session *sessionService) Close() {
	log.Printf("%s: service is closed\n", session.TypeName())
}

func (session *sessionService) TypeName() string {
	return "session"
}
