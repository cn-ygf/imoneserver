package business

import (
	"github.com/cn-ygf/yin"
	"sync"
	"github.com/cn-ygf/imoneserver/lib/crypto"
)

var(
	vfsessions = sync.Map{}
)

// 根据vfsession取得vfcode
func GetVfCode(vfsession string)(vfcode string){
	if v, ok := vfsessions.Load(vfsession); ok {
		vfcode = v.(string)
	}
	vfsessions.Delete(vfsession)
	return
}

// 获取随机验证码
func VfCode(ctx yin.Context) {
	vfcode := crypto.GetRandomString(6)
	vfsession := crypto.Md5(crypto.GetRandomString(32))
	vfsessions.Store(vfsession,vfcode)
	ctx.SUCCESS(map[string]interface{}{
		"vfcode": vfcode,
		"vfsession": vfsession,
	})
}