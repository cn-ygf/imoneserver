package business

import (
	"github.com/cn-ygf/imoneserver/api/dao/member"
	"github.com/cn-ygf/imoneserver/lib/database/orm"
	"github.com/cn-ygf/yin"
)

var (
	d *member.Dao
)

func init() {
	d = member.New(&orm.Config{
		DSN: "root:root@/im_one?charset=utf8&parseTime=True&loc=Local",
	})
}

// 用户登录
func Login(ctx yin.Context) {
	ctx.HTML(200, "test")
}

func Member(ctx yin.Context) {
	res, err := d.Members()
	if err != nil {
		ctx.HTML(200, err.Error())
		return
	}
	ctx.JSON(200, res)
}
