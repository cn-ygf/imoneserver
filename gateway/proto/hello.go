// 首次握手
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
		Type:  reflect.TypeOf((*HelloREQ)(nil)).Elem(),
		ID:    int(MSG_HELLO_REQ),
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("protobuf"),
		Type:  reflect.TypeOf((*HelloACK)(nil)).Elem(),
		ID:    int(MSG_HELLO_ACK),
	})
}
