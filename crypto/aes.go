/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

//aes加密
func AesEncryptSimple(origData string, key string, iv string) (string, error) {
	xpass, err := AesEncryptPkcs5([]byte(origData), []byte(key), []byte(iv))
	if err != nil {
		return "", err
	}
	pass64 := base64.StdEncoding.EncodeToString(xpass)
	return pass64, nil
}

//aes解密
func AesDecryptSimple(crypted string, key string, iv string) (string, error) {
	bytesPass, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		return "", err
	}

	res, err := AesDecryptPkcs5(bytesPass, []byte(key), []byte(iv))
	return string(res), err
}

func AesEncryptPkcs5(origData []byte, key []byte, iv []byte) ([]byte, error) {
	return AesEncrypt(origData, key, iv, PKCS5Padding)
}

func AesDecryptPkcs5(crypted []byte, key []byte, iv []byte) ([]byte, error) {
	return AesDecrypt(crypted, key, iv, PKCS5UnPadding)
}

func AesEncrypt(origData []byte, key []byte, iv []byte, paddingFunc func([]byte, int) []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = paddingFunc(origData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte, iv []byte, unPaddingFunc func([]byte) []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = unPaddingFunc(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	if length < unpadding {
		return []byte("unpadding error")
	}
	return origData[:(length - unpadding)]
}
