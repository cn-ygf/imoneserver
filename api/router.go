package api

import (
	"github.com/cn-ygf/imoneserver/api/business"
	"github.com/cn-ygf/yin"
)

// 注册路由
func RegisterRouter(engine yin.Engine) {
	engine.GET("/api/v1/login", business.Login)
}
