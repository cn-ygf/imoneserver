package member

import (
	"github.com/cn-ygf/imoneserver/api/model/member"
	"github.com/cn-ygf/imoneserver/lib/database/orm"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	db *gorm.DB
}

func New(c *orm.Config) (d *Dao) {
	d = &Dao{
		db: orm.NewMySQL(c),
	}
	d.initORM()
	return
}

func (d *Dao) initORM() {
	d.db.LogMode(true)
}

// 获取所有用户
func (d *Dao) Members() (res []*member.Member, err error) {
	if err := d.db.Order("id DESC", true).Find(&res).Error; err != nil {
		log.Errorln("d.db.Order err(%v)", err)
		return nil, err
	}
	return
}

// 根据邮箱获取用户
func (d *Dao) MemberByEmail(email string) (res *member.Member, err error) {
	res = &member.Member{}
	if err = d.db.Where("email = ?", email).First(res).Error; err != nil {
		log.Errorln("d.db.Where error(%v)", err)
		return nil, err
	}
	return
}
