package session

import (
	"github.com/cn-ygf/imoneserver"
	"github.com/cn-ygf/imoneserver/service"
)

// 注册服务
func init() {
	service.RegisterServiceCreator(func() imoneserver.Service {
		s := &sessionService{}
		return s
	})
}
