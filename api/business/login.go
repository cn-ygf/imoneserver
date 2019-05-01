package business

import "github.com/cn-ygf/yin"

// 用户登录
func Login(ctx yin.Context) {
	ctx.HTML(200, "test")
}
