package api

import (
	"fmt"
	"github.com/cn-ygf/imoneserver/api/business"
	"github.com/cn-ygf/yin"
)

// 注册路由
func RegisterRouter(engine yin.Engine, v string) {
	engine.GET(fmt.Sprintf("/api/%s/vfcode", v), business.VfCode)
	engine.POST(fmt.Sprintf("/api/%s/login", v), business.Login)
	engine.GET(fmt.Sprintf("/api/%s/images", v), business.GetImages)
	engine.GET(fmt.Sprintf("/api/%s/member/info", v), business.MemberInfo)
	engine.GET("/api/test/member", business.Member)
}
