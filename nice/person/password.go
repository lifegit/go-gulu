/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package person

import (
	"errors"
	"github.com/lifegit/go-gulu/v2/nice/crypto"
	"github.com/lifegit/go-gulu/v2/nice/rand"
	"github.com/wenzhenxi/gorsa"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

type Password struct {
	Code        string // code 随即码
	Salt        string // 盐
	SaltEncrypt string // 加密盐
}

// 生产salt和code
func (a *Password) RandSaltAndCode(privateKey string) error {
	a.Salt = rand.String(rand.SeedAll, 16)
	a.Code = rand.String(rand.SeedNum, rand.Int(6, 12))
	saltEncrypt, err := gorsa.PriKeyEncrypt(a.Salt, privateKey)
	a.SaltEncrypt = saltEncrypt
	return err
}

/**
  * 解密密码
  * 原理:
     //     AES-128-CBC {padding:Pkcs7,iv:salt}(
     //         md5( md5(salt + 验证码)+sha512(单数位 salt) ).从第三个开始取16个长度
     //     AES-128-CBC {padding:Pkcs7,iv:md5(salt)的前16位}(
     //         sha512(sha384(salt) + 验证码长度 + 验证码 + md5(双数位 salt)).从第十个开始取16个长度,
     //         sha512(sha384(salt) + 验证码长度 + 验证码 + md5(双数位 salt)) + QQ 密码)
     //     )
     // 注: salt 长度必须为16
*/
// 解密密码与验证
func (a *Password) DecryptAndCheck(hashedPassword string, inputPass string) (bool, error) {
	// step1: decrypt pass
	dSalt := ""
	sSalt := ""
	// 计算单双
	for i := 0; i < len(a.Salt); i++ {
		if i%2 == 0 {
			dSalt += string(a.Salt[i])
		} else {
			sSalt += string(a.Salt[i])
		}
	}

	// 计算key
	keyKey := crypto.EncodeMD5(crypto.EncodeMD5(a.Salt+a.Code) + crypto.EncodeSha512(dSalt))[3:19] //19 = 3+16
	keyValue := crypto.EncodeSha512(crypto.EncodeSha384(a.Salt) + strconv.Itoa(len(a.Code)) + a.Code + crypto.EncodeMD5(sSalt))

	// 密码解密
	decryptPassword, err := crypto.AesDecryptSimple(inputPass, keyKey, a.Salt)
	if err != nil {
		return false, err
	}
	decryptPassword, err = crypto.AesDecryptSimple(decryptPassword, keyValue[10:26], crypto.EncodeMD5(a.Salt)[0:16]) //26 = 10 + 16
	if err != nil {
		return false, err
	}

	decryptPassword = strings.Replace(decryptPassword, keyValue, "", 1)

	// step2: password is set to bcrypt check
	return a.Check(hashedPassword, decryptPassword)
}

// 验证一个hash后的密码
func (a *Password) Check(hashedPassword, password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false, err
	}
	return true, nil
}

// 加密密码
func (*Password) MakePassword(password string) (bcryptPass string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("bcrypt making password is failed")
	}

	return string(bytes), nil
}
