package tcp

import (
	"io"
	"net"

	"github.com/cn-ygf/imoneserver/gateway/session"
	"github.com/davyxu/cellnet"
)

type TCPMessageTransmitter struct {
	Session *session.Session
}

func NewTCPMessageTransmitter() *TCPMessageTransmitter {
	t := new(TCPMessageTransmitter)
	t.Session = nil
	return t
}

type socketOpt interface {
	MaxPacketSize() int
	ApplySocketReadTimeout(conn net.Conn, callback func())
	ApplySocketWriteTimeout(conn net.Conn, callback func())
}

func (TCPMessageTransmitter) OnRecvMessage(ses cellnet.Session) (msg interface{}, err error) {

	reader, ok := ses.Raw().(io.Reader)

	// 转换错误，或者连接已经关闭时退出
	if !ok || reader == nil {
		return nil, nil
	}

	opt := ses.Peer().(socketOpt)

	if conn, ok := reader.(net.Conn); ok {

		// 有读超时时，设置超时
		opt.ApplySocketReadTimeout(conn, func() {
			msg, err = RecvYsocksPacket(reader, ses /*ses.Peer().(cellnet.ContextSet)*/, opt.MaxPacketSize())

		})
	}

	return
}

func (TCPMessageTransmitter) OnSendMessage(ses cellnet.Session, msg interface{}) (err error) {

	writer, ok := ses.Raw().(io.Writer)

	// 转换错误，或者连接已经关闭时退出
	if !ok || writer == nil {
		return nil
	}

	opt := ses.Peer().(socketOpt)

	// 有写超时时，设置超时
	opt.ApplySocketWriteTimeout(writer.(net.Conn), func() {
		err = SendYsocksPacket(writer, ses, msg)
	})

	return
}
