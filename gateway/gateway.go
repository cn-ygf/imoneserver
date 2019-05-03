package gateway

import (
	_ "github.com/cn-ygf/imoneserver/gateway/proc/tcp"
	"github.com/cn-ygf/imoneserver/gateway/proto"
	"github.com/cn-ygf/imoneserver/gateway/session"
	"github.com/cn-ygf/imoneserver/lib/crypto"
	"github.com/cn-ygf/imoneserver/lib/crypto/server"
	"github.com/cn-ygf/imoneserver/service"
	sess "github.com/cn-ygf/imoneserver/session"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/tcp"
	"github.com/davyxu/cellnet/proc"
	pb "github.com/golang/protobuf/proto"
	"time"
)

type gateWayService struct {
	queue cellnet.EventQueue
	p     cellnet.Peer
	name  string
}

func (gate *gateWayService) Run(param ...interface{}) {
	// 创建服务器密钥
	server.Init()

	localAddress := "0.0.0.0:9008"
	if len(param) > 0 {
		localAddress = param[0].(string)
	}
	gate.queue = cellnet.NewEventQueue()
	gate.p = peer.NewGenericPeer("tcp.Acceptor", "server", localAddress, gate.queue)
	proc.BindProcessorHandler(gate.p, "tcp.imone", gate.handler)
	gate.p.Start()
	gate.queue.StartLoop()
	log.Infof("%s: service is running", gate.Name())
}

func (gate *gateWayService) Close() {
	gate.p.Stop()
	gate.queue.StopLoop()
	log.Infof("%s: service is closed", gate.Name())
}

func (gate *gateWayService) TypeName() string {
	return "gateway"
}

func (gate *gateWayService) handler(ev cellnet.Event) {
	switch msg := ev.Message().(type) {
	case *cellnet.SessionAccepted:
		log.Debugln("server accepted:id ", ev.Session().ID())
		log.Debugln(msg.String())
	case *cellnet.SessionClosed:
		log.Debugln("session closed: ", ev.Session().ID())
	case *proto.HelloREQ:
		log.Debugln("hello req:", ev.Session().ID())
		gate.helloAck(ev, msg)
	case *proto.LoginREQ:
		log.Debugln("login req:", ev.Session().ID())
		gate.loginAck(ev, msg)
	case *proto.HBPREQ:
	}
}

func (gate *gateWayService) Name() string {
	if len(gate.name) < 1 {
		return gate.TypeName()
	}
	return gate.name
}

func (gate *gateWayService) SetName(name string) {
	gate.name = name
}

// Hello 应答
func (gate *gateWayService) helloAck(ev cellnet.Event, msg *proto.HelloREQ) {
	timecom := time.Now().Unix() - msg.GetTime()
	if timecom < 0 {
		timecom = ^timecom + 1
	}
	if timecom > 30 {
		ev.Session().Close()
		log.Debugln("hello req time failed:", ev.Session().ID())
		return
	}
	if msg.GetMsg() != "hello imone" {
		ev.Session().Close()
		log.Debugln("hello req failed:", ev.Session().ID())
		return
	}
	// 发送服务器公钥给客户端
	ack := &proto.HelloACK{
		Code: pb.Int32(CodeSuccess),
		Msg:  pb.String("hello"),
		Pub:  server.GetPub(),
	}
	ev.Session().Send(ack)
}

// login应答
func (gate *gateWayService) loginAck(ev cellnet.Event, msg *proto.LoginREQ) {
	timecom := time.Now().Unix() - msg.GetTime()
	if timecom < 0 {
		timecom = ^timecom + 1
	}
	if timecom > 30 {
		ev.Session().Close()
		log.Debugln("login req time failed:", ev.Session().ID())
		return
	}
	// 解密密码
	passBuffer, err := crypto.RSADecrypt(server.GetPrv(), msg.GetPassword())
	if err != nil {
		log.Debugln("login req rsa decrypt failed!error:", err.Error())
		ev.Session().Close()
		return
	}
	password := string(passBuffer)
	sessionService := service.GetService("session")
	sessObj := sessionService.(sess.SessionMgr).Get(password)
	// 判断是否存在
	if sessObj == nil {
		log.Debugln("login req failed!error:", "password error")
		ev.Session().Close()
		return
	}
	clientSession := sessObj.EvSession()
	if clientSession != nil {
		// TODO 强制下线
		clientSession.Close()
	}
	// 绑定会话
	sessObj.SetEvSession(ev.Session())
	// 创建Session
	token := session.New()
	ev.Session().(cellnet.ContextSet).SetContext("session", token)
	// 组建回应包
	loginACK := &proto.LoginACK{
		Code:       pb.Int32(CodeSuccess),
		Sessionkey: token.Key(),
	}
	// 发送
	ev.Session().Send(loginACK)
	log.Debugln("session create:", ev.Session().ID())
}
