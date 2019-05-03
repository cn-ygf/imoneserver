package session

import "github.com/cn-ygf/imoneserver/api/model/member"

type SessionMgr interface {
	New(*member.Member) string // 创建session
	Remove(key string)         // 删除session
	Get(key string) Session    // 获取session
}
