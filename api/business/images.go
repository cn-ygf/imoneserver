package business

import (
	"encoding/hex"
	"fmt"
	"github.com/cn-ygf/imoneserver/lib/config"
	"github.com/cn-ygf/yin"
)

// 获取图片
func GetImages(ctx yin.Context) {
	hash := ctx.Body().Get("hash")
	if len(hash) != 32 {
		ctx.ERROR(nil)
		return
	}
	// 过虑非hash的参数
	_, err := hex.DecodeString(hash)
	if err != nil {
		log.Errorln(err)
		ctx.ERROR(nil)
		return
	}
	path := fmt.Sprintf("%s/%s", config.GetString("storage_dir"), hash)
	ctx.FILE(200, path, "image/png")
}
