package protobuf

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	"github.com/golang/protobuf/proto"
)

// protobuf解码器
type protobufCodec struct {
}

// 编码器名称
func (protobufCodec) Name() string {
	return "protobuf"
}

// 属性
func (protobufCodec) MimeType() string {
	return "application/protobuf"
}

// 编码数据
func (protobufCodec) Encode(msgObj interface{}, ctx cellnet.ContextSet) (data interface{}, err error) {
	return proto.Marshal(msgObj.(proto.Message))

}

// 解码数据
func (protobufCodec) Decode(data interface{}, msgObj interface{}) error {
	return proto.Unmarshal(data.([]byte), msgObj.(proto.Message))
}

func init() {

	// 注册编码器
	codec.RegisterCodec(new(protobufCodec))
}
