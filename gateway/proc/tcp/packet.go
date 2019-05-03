package tcp

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"io"

	"github.com/cn-ygf/imoneserver/gateway/proto"

	"github.com/cn-ygf/imoneserver/gateway/session"
	"github.com/cn-ygf/imoneserver/lib/crypto"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
)

var (
	ErrMaxPacket     = errors.New("packet over size")
	ErrMinPacket     = errors.New("packet short size")
	ErrShortMsgID    = errors.New("short msgid")
	ErrHeadPacket    = errors.New("error head")
	ErrVersionPacket = errors.New("error version")
	ErrFooterPacket  = errors.New("error footer")
	ErrCryptoPacket  = errors.New("error iscrypto")
	ErrConverStruct  = errors.New("error converstruct")
)

const (
	bodySize  = 2 // 包体大小字段
	msgIDSize = 2 // 消息ID字段
	headSize  = 8 // 包头大小
)

// 接收ysocks格式的封包流程
func RecvYsocksPacket(reader io.Reader, ses cellnet.Session, maxPacketSize int) (msg interface{}, err error) {
	// 创建head缓存，包头为8个字节
	var headBuffer = make([]byte, headSize)
	// 持续读取headSize
	_, err = io.ReadFull(reader, headBuffer)
	if err != nil {
		return
	}
	if len(headBuffer) < bodySize {
		return nil, ErrMinPacket
	}
	// 验证包头
	msgId, err := HeadVerify(headBuffer[0:7])
	if err != nil {
		return
	}
	// 获取包长度
	packageSize := binary.LittleEndian.Uint16(headBuffer[5:])
	// 判断包长度是否合法
	if maxPacketSize > 0 && packageSize >= uint16(maxPacketSize) {
		return nil, ErrMaxPacket
	}
	// 判断是否需要解密
	if headBuffer[7] != 0x00 && headBuffer[7] != 0x01 {
		return nil, ErrCryptoPacket
	}
	// 分配包体大小
	body := make([]byte, packageSize+1)
	// 读取包体数据
	_, err = io.ReadFull(reader, body)
	// 发生错误时返回
	if err != nil {
		return
	}
	// 验证包尾
	if body[len(body)-1] != 0x03 {
		log.Debugln("head buf :", headBuffer)
		log.Debugln("body len :", len(body), body)
		return nil, ErrFooterPacket
	}
	// 需求处理的包
	packageBuffer := body[0:packageSize]
	// 判断是否需要解密
	if headBuffer[7] == 0x01 {
		s, ok := ses.(cellnet.ContextSet).GetContext("session")
		if s == nil {
			// TODO 解密失败
			return nil, ErrCryptoPacket
		}
		sess, ok := s.(session.Session)
		if !ok {
			// TODO 解密失败
			return nil, ErrConverStruct
		}
		// 解密
		packageBuffer, err = crypto.AesDecrypt(packageBuffer, sess.Key(), sess.Key())
		if err != nil {
			// TODO 解密失败
			return nil, err
		}
		// 解压
		packageBuffer = DoZlibUnCompress(packageBuffer)
	}
	// 将字节数组和消息ID用户解出消息
	msg, _, err = codec.DecodeMessage(int(msgId), packageBuffer)
	if err != nil {
		// TODO 接收错误时，返回消息
		return nil, err
	}
	return
}

// 发送ysocks格式的封包流程
func SendYsocksPacket(writer io.Writer, ses cellnet.Session, data interface{}) error {
	var (
		msgData  []byte
		msgID    uint16
		meta     *cellnet.MessageMeta
		isCrypto byte
		sess     session.Session
	)
	ctx := ses.(cellnet.ContextSet)
	// 默认不需要加密
	isCrypto = 0x00
	s, ok := ctx.GetContext("session")
	if ok && s != nil {
		sess = s.(session.Session)
		isCrypto = 0x01
	}
	switch m := data.(type) {
	case *cellnet.RawPacket: // 发裸包
		msgData = m.MsgData
		msgID = uint16(m.MsgID)
	case *proto.LoginACK:
		isCrypto = 0x00
		var err error
		// 将用户数据转换为字节数组和消息ID
		msgData, meta, err = codec.EncodeMessage(data, ctx)
		if err != nil {
			return err
		}
		msgID = uint16(meta.ID)
	default: // 发普通编码包
		var err error

		// 将用户数据转换为字节数组和消息ID
		msgData, meta, err = codec.EncodeMessage(data, ctx)

		if err != nil {
			return err
		}

		msgID = uint16(meta.ID)
	}
	// 需求处理的包
	packageBuffer := msgData
	// 判断是否需要加密
	if isCrypto == 0x01 {
		var err error
		// 压缩
		packageBuffer = DoZlibCompress(packageBuffer)
		// 加密
		packageBuffer, err = crypto.AesEncrypt(packageBuffer, sess.Key(), sess.Key())
		if err != nil {
			return err
		}
	}

	// 组建发送数据包
	sendBytes := bytes.NewBuffer(nil)
	// 写入包头
	sendBytes.WriteByte(0x02)
	// 写入版本号
	ver := uint16(0x01)
	binary.Write(sendBytes, binary.LittleEndian, ver)
	// 写入消息id
	binary.Write(sendBytes, binary.LittleEndian, msgID)
	// 写入body长度
	bodySize := uint16(len(packageBuffer))

	binary.Write(sendBytes, binary.LittleEndian, bodySize)
	// 写入是否需要加密
	sendBytes.WriteByte(isCrypto)
	// 写入body
	sendBytes.Write(packageBuffer)
	// 写入包尾
	sendBytes.WriteByte(0x03)
	// 将数据写入Socket
	err := WriteFull(writer, sendBytes.Bytes())

	// Codec中使用内存池时的释放位置
	if meta != nil {
		codec.FreeCodecResource(meta.Codec, msgData, ctx)
	}

	return err
}

// 完整发送所有封包
func WriteFull(writer io.Writer, buf []byte) error {

	total := len(buf)

	for pos := 0; pos < total; {

		n, err := writer.Write(buf[pos:])

		if err != nil {
			return err
		}

		pos += n
	}

	return nil

}

// 验证包头
func HeadVerify(headBuffer []byte) (uint16, error) {
	// 包头
	if headBuffer[0] != 0x02 {
		return 0, ErrHeadPacket
	}
	// 版本号
	ver := binary.LittleEndian.Uint16(headBuffer[1:])
	if ver != 0x01 {
		return 0, ErrVersionPacket
	}
	// 消息id
	msgId := binary.LittleEndian.Uint16(headBuffer[3:])
	return msgId, nil
}

// 进行zlib压缩
func DoZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

// 进行zlib解压缩
func DoZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}
