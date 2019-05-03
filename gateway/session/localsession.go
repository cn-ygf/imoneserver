// 与客户端之间的会话
package session

type localSession struct {
	sessionKey []byte // Session key 会话密钥
	created    int64  // 创建时间
}

func (local *localSession) Key() []byte {
	return local.sessionKey
}

func (local *localSession) Created() int64 {
	return local.created
}
