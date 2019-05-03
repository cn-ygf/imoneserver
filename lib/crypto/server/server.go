// 保存服务器密钥对
package server

import (
	"github.com/cn-ygf/imoneserver/lib/crypto"
	"github.com/davyxu/golog"
)

var log = golog.New("crypto.server")

var pub []byte
var prv []byte

func Init() {
	var err error
	prv, pub, err = crypto.RSAKeyGen()
	if err != nil {
		log.Errorln(err.Error())
		return
	}
}

// 获取公钥
func GetPub() []byte {
	return pub
}

// 获取私钥
func GetPrv() []byte {
	return prv
}
