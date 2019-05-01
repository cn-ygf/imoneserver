// hash相关函数
package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// 计算md5值
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	c := h.Sum(nil)
	return hex.EncodeToString(c)
}

// 计算md5值byte
func Md5Bytes(str string) []byte {
	h := md5.New()
	h.Write([]byte(str))
	c := h.Sum(nil)
	return c
}

// 计算sha256
func Sha256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	c := h.Sum(nil)
	return hex.EncodeToString(c)
}
