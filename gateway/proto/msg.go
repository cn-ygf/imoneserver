// 消息id
package proto

const (
	MSG_HELLO_REQ = 10001 // 首次握手
	MSG_HELLO_ACK = 10002 // 首次握手服务器回应
	MSG_LOGIN_REQ = 10003 // 登录请求
	MSG_LOGIN_ACK = 10004 // 登录请求服务器回应
	MSG_HBP_REQ   = 10005 // 心跳包请求
	MSG_HBP_ACK   = 10006 // 心跳包服务器回应
)
