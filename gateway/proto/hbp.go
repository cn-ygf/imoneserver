// 心跳包消息
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
		Type:  reflect.TypeOf((*HBPREQ)(nil)).Elem(),
		ID:    int(MSG_HBP_REQ),
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("protobuf"),
		Type:  reflect.TypeOf((*HBPACK)(nil)).Elem(),
		ID:    int(MSG_HBP_ACK),
	})
}
