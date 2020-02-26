/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package crypto

import (
	"crypto/rc4"
	"encoding/base64"
)

// 加密
func Rc4EncryptSimple(key string, origData string) (text string) {
	res := Rc4Encrypt([]byte(key), []byte(origData))
	return base64.StdEncoding.EncodeToString(res)
}

// 加密
func Rc4Encrypt(key []byte, origData []byte) (text []byte) {
	c, err := rc4.NewCipher(key)
	if err != nil {
		return
	}
	dst := make([]byte, len(origData))
	c.XORKeyStream(dst, origData)

	return dst
}

// 解密
func Rc4DecryptSimple(key string, crypted string) (text string) {
	src, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		return
	}
	return string(Rc4Decrypt([]byte(key), src))
}

// 解密
func Rc4Decrypt(key []byte, crypted []byte) (text []byte) {
	c, err := rc4.NewCipher(key)
	if err != nil {
		return
	}
	dst := make([]byte, len(crypted))
	c.XORKeyStream(dst, crypted)

	return dst
}
