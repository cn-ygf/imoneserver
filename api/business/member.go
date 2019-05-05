package business

import (
	"fmt"
	"github.com/cn-ygf/imoneserver/api/dao/member"
	member2 "github.com/cn-ygf/imoneserver/api/model/member"
	"github.com/cn-ygf/imoneserver/lib/config"
	"github.com/cn-ygf/imoneserver/lib/crypto"
	"github.com/cn-ygf/imoneserver/lib/database/orm"
	"github.com/cn-ygf/imoneserver/service"
	"github.com/cn-ygf/imoneserver/session"
	"github.com/cn-ygf/yin"
)

var (
	d *member.Dao
)

func Init() {
	d = member.New(&orm.Config{
		DSN: config.GetString("mysql_dsn"),
	})
}

// 用户登录
func Login(ctx yin.Context) {
	email := ctx.Body().Get("email")
	password := ctx.Body().Get("password")
	vfcode := ctx.Body().Get("vfcode")
	vfsession := ctx.Body().Get("vfsession")
	m, err := d.MemberByEmail(email)
	if err != nil {
		log.Errorln(err)
		ctx.ERROR(nil)
		return
	}
	if m.Id == 0 {
		ctx.SUCCESS(map[string]interface{}{
			"code": 10001,
			"msg":  "email or password error",
		})
		return
	}
	trueVfCode := GetVfCode(vfsession)
	if len(trueVfCode) < 1 {
		ctx.SUCCESS(map[string]interface{}{
			"code": 10002,
			"msg":  "vfcode error",
		})
		return
	}
	if vfcode != trueVfCode {
		ctx.SUCCESS(map[string]interface{}{
			"code": 10002,
			"msg":  "vfcode error",
		})
		return
	}
	truePassword := crypto.Md5(fmt.Sprintf("%simone2019%s", m.Password, trueVfCode))
	if password != truePassword {
		ctx.SUCCESS(map[string]interface{}{
			"code": 10001,
			"msg":  "email or password error",
		})
		return
	}
	sessionKey := service.GetService("session").(session.SessionMgr).New(m)
	ctx.SUCCESS(map[string]interface{}{
		"code":       10000,
		"msg":        "success",
		"sessionkey": sessionKey, //TODO
	})
	log.Debugln("sessionkey:", sessionKey)
}

func Member(ctx yin.Context) {
	res, err := d.Members()
	if err != nil {
		ctx.HTML(200, err.Error())
		return
	}
	ctx.JSON(200, res)
}

// 获取用户信息
func MemberInfo(ctx yin.Context) {
	sessionKey := ctx.Body().Get("sessionkey")
	if len(sessionKey) != 32 {
		ctx.ERROR(nil)
		return
	}
	sess := service.GetService("session").(session.SessionMgr).Get(sessionKey)
	if sess == nil {
		// session过期或不存在
		ctx.ERROR(map[string]interface{}{
			"code": 10001,
			"msg":  "session error",
		})
		return
	}
	// 取得member对象
	m := sess.Object().(*member2.Member)
	ctx.SUCCESS(map[string]interface{}{
		"code": 10000,
		"msg":  "success",
		"member": map[string]interface{}{
			"id":      m.Id,
			"email":   m.Email,
			"name":    m.Name,
			"nick":    m.Nick,
			"head":    m.Head,
			"current": m.Current,
			"phone":   m.Phone,
		},
	})
}

// 获取联系人
func Contacts(ctx yin.Context) {
	// TODO
}
