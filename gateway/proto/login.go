// 登录消息
package proto

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"

	"reflect"
)

// 注册消息
func init() {
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("protobuf"),
		Type:  reflect.TypeOf((*LoginREQ)(nil)).Elem(),
		ID:    int(MSG_LOGIN_REQ),
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("protobuf"),
		Type:  reflect.TypeOf((*LoginACK)(nil)).Elem(),
		ID:    int(MSG_LOGIN_ACK),
	})
}
