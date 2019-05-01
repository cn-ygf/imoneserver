package member

import (
	"github.com/cn-ygf/imoneserver/api/model/member"
	"github.com/cn-ygf/imoneserver/lib/database/orm"
	"github.com/jinzhu/gorm"
	"log"
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
		log.Printf("d.db.Order err(%v)", err)
		return nil, err
	}
	return
}
