package member

import (
	"github.com/cn-ygf/imoneserver/lib/database/orm"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	d *Dao
)

func init() {
	d = New(&orm.Config{
		DSN: "root:root@/im_one?charset=utf8&parseTime=True&loc=Local",
	})
}

func TestMemberByEmail(t *testing.T) {
	Convey("select by email", t, func() {
		res, err := d.MemberByEmail("ygf@cnhonker.com")
		So(err, ShouldBeNil)
		So(res, ShouldNotBeEmpty)
	})
}

