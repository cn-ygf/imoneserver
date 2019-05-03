// 会话管理
package session

import (
	"github.com/cn-ygf/imoneserver/lib/crypto"
	"time"
)

type Session interface {
	Key() []byte    // 取得会话密钥
	Created() int64 // 获取创建时间
}

// 创建会话
func New() Session {
	randStr := crypto.GetRandomString(64)
	sessionKey := crypto.Md5Bytes(randStr)
	sess := &localSession{
		sessionKey: sessionKey,
		created:    time.Now().Unix(),
	}
	return sess
}

// 绑定会话
func Bind(sessionKey []byte) Session {
	sess := &localSession{
		sessionKey: sessionKey,
		created:    time.Now().Unix(),
	}
	return sess
}
