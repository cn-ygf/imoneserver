package business

import (
	"fmt"
	"github.com/cn-ygf/imoneserver/api/dao/member"
	"github.com/cn-ygf/imoneserver/lib/crypto"
	"github.com/cn-ygf/imoneserver/lib/database/orm"
	"github.com/cn-ygf/imoneserver/service"
	"github.com/cn-ygf/imoneserver/session"
	"github.com/cn-ygf/yin"
	"log"
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
	email:= ctx.Body().Get("email")
	password := ctx.Body().Get("password")
	vfcode := ctx.Body().Get("vfcode")
	vfsession := ctx.Body().Get("vfsession")
	m,err := d.MemberByEmail(email)
	if err != nil{
		log.Println(err)
		ctx.ERROR(nil)
		return
	}
	if m.Id  == 0{
		ctx.SUCCESS(map[string]interface{}{
			"code":10001,
			"msg":"email or password error",
		})
		return
	}
	trueVfCode := GetVfCode(vfsession)
	if len(trueVfCode) < 1 {
		ctx.SUCCESS(map[string]interface{}{
			"code":10002,
			"msg":"vfcode error",
		})
		return
	}
	if vfcode != trueVfCode {
		ctx.SUCCESS(map[string]interface{}{
			"code":10002,
			"msg":"vfcode error",
		})
		return
	}
	truePassword := crypto.Md5(fmt.Sprintf("%simone2019%s",m.Password,trueVfCode))
	if password != truePassword{
		ctx.SUCCESS(map[string]interface{}{
			"code":10001,
			"msg":"email or password error",
		})
		return
	}
	sessionKey := service.GetService("session").(session.SessionMgr).New(m)
	ctx.SUCCESS(map[string]interface{}{
		"code":10000,
		"msg":"success",
		"sessionkey":sessionKey,//TODO
	})
}

func Member(ctx yin.Context) {
	res, err := d.Members()
	if err != nil {
		ctx.HTML(200, err.Error())
		return
	}
	ctx.JSON(200, res)
}
