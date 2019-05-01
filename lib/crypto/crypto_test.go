package crypto

import (
	"testing"
)

func TestRSAKeyGen(t *testing.T) {
	prv, pub, err := RSAKeyGen()
	if err != nil {
		t.Error(err.Error())
		return
	}
	//prvHex := hex.EncodeToString(prv)
	//pubHex := hex.EncodeToString(pub)
	//t.Log(prvHex)
	//t.Log(pubHex)
	mi := []byte("hello ysocks!")
	mi2, err := RSAEncrypt(pub, mi)
	if err != nil {
		t.Error(err.Error())
		return
	}
	mi3, err := RSADecrypt(prv, mi2)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(string(mi))
	t.Log(mi2)
	t.Log(string(mi3))
}
