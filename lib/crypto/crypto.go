// 加密算法
package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 生成rsa密钥对
func RSAKeyGen() ([]byte, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, err
	}
	prvStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA Private key",
		Bytes: prvStream,
	}
	prv := pem.EncodeToMemory(block)
	publicKey := &privateKey.PublicKey
	pubSteam, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}
	block = &pem.Block{
		Type:  "RSA Public key",
		Bytes: pubSteam,
	}
	pub := pem.EncodeToMemory(block)
	return prv, pub, nil
}

// rsa加密
func RSAEncrypt(pubKey []byte, data []byte) ([]byte, error) {
	// 解码公钥
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return nil, errors.New("pubkey error")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), data)
}

// rsa解密
func RSADecrypt(prvKey []byte, data []byte) ([]byte, error) {
	// 解码公钥
	block, _ := pem.Decode(prvKey)
	if block == nil {
		return nil, errors.New("prvkey error")
	}
	prv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, prv, data)
}

// aes加密
// origData:待加密的数据
// key:密钥
// iv:aes用到的iv
func AesEncrypt(origData []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// aes解密
// decodeBytes:待解密的数据
// key:密钥
// iv:aes用到的iv
func AesDecrypt(decodeBytes []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
