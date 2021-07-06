/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package crypto

import (
	"github.com/speps/go-hashids"
)

func initHashids(salt string, length int) (*hashids.HashID) {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = length
	return hashids.NewWithData(hd)
}

//// 使用salt对str进行编码
//func EncodeHex(salt string, length int, str string) (string, error) {
//	res, err := initHashids(salt, length).EncodeHex(str)
//	if err != nil {
//		return "", err
//	}
//
//	return res, nil
//}
//
//// 使用salt对str进行解码
//func DecodeHex(salt string, length int, str string) (string, error) {
//	hashID, err := initHashids(salt, length)
//	if err != nil {
//		return "", err
//	}
//
//	res, err := hashID.DecodeHex(str)
//	if err != nil {
//		return "", err
//	}
//
//	return res, nil
//}
