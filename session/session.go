package session

import (
	"github.com/cn-ygf/imoneserver/api/model/member"
	"github.com/cn-ygf/imoneserver/lib/crypto"
	"github.com/davyxu/cellnet"
	"sync"
	"time"
)

// 会话管理服务
type sessionService struct {
	sessions sync.Map
}

func (session *sessionService) Run(param ...interface{}) {
	log.Infof("%s: service is running\n", session.TypeName())
}

func (session *sessionService) Close() {
	log.Infof("%s: service is closed\n", session.TypeName())
}

func (session *sessionService) TypeName() string {
	return "session"
}

func (session *sessionService) New(m *member.Member) string {
	sessionKey := crypto.Md5(crypto.GetRandomString(64))
	sess := &coreSession{
		sessionKey: sessionKey,
		obj:        m,
		date:       time.Now().Unix(),
		device:     1,
	}
	session.sessions.Store(sessionKey, sess)
	return sessionKey
}

func (session *sessionService) Remove(key string) {
	session.sessions.Delete(key)
}

func (session *sessionService) Get(key string) Session {
	if v, ok := session.sessions.Load(key); ok {
		return v.(Session)
	}
	return nil
}

type Session interface {
	String() string               // 取得sessionkey
	Object() interface{}          // 绑定的member对象
	Date() int64                  // 登录时间
	Device() int64                // 登录设备类型
	LastDate() int64              // 最后一次连接时间
	EvSession() cellnet.Session   // 连接成功的会话
	SetEvSession(cellnet.Session) // 关联会话
}

type coreSession struct {
	sessionKey string
	obj        *member.Member
	date       int64
	device     int64
	lastDate   int64
	evsess     cellnet.Session
}

func (core *coreSession) String() string {
	return core.sessionKey
}

func (core *coreSession) Object() interface{} {
	return core.obj
}
func (core *coreSession) Date() int64 {
	return core.date
}
func (core *coreSession) Device() int64 {
	return core.device
}
func (core *coreSession) LastDate() int64 {
	return core.lastDate
}

func (core *coreSession) EvSession() cellnet.Session {
	return core.evsess
}
func (core *coreSession) SetEvSession(s cellnet.Session) {
	core.evsess = s
}
