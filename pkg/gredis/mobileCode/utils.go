/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package mobileCode

import (
	"github.com/lifegit/go-gulu/v2/nice/rand"
	"regexp"
)

func tag(tag, mobile string) string {
	return tag + mobile
}

// 验证一个手机号
func MobileValidate(mobile string) bool {
	match, _ := regexp.MatchString(`^[1][3,4,5,6,7,8,9][0-9]{9}$`, mobile)
	return match
}

// 手机号隐私化
func MobilePrivate(mobile string) (b bool, text string) {
	if len(mobile) != 11 {
		return false, mobile
	}

	return true, mobile[0:3] + "*****" + mobile[8:]
}

// 获取随机验证码
func RandCode(len int) string {
	return rand.String(rand.SeedNum, len)
}
