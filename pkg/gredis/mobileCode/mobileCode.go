/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package mobileCode

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type MobileCode struct {
	tag    string
	expire time.Duration
	redis  *redis.Client
}

func New(tag string, redis *redis.Client, expire time.Duration) *MobileCode {
	return &MobileCode{
		tag:    fmt.Sprintf("%s:", tag),
		expire: expire,
		redis:  redis,
	}
}
// 发送
func (m *MobileCode) Send(sendFunc func() (MobileMes, error)) error {
	// 发送验证码
	sendMes, err := sendFunc()
	if err != nil {
		return err
	}

	// 放到缓存
	return m.redis.Set(tag(m.tag, sendMes.Mobile), sendMes.Code, m.expire).Err()
}
// 是否正确
func (m *MobileCode) IsCheck(ms MobileMes) bool {
	str, _ := m.redis.Get(tag(m.tag, ms.Mobile)).Result()
	return str == ms.Code
}
// 是否存在
func (m *MobileCode) IsExist(mobile string) bool {
	str, _ := m.redis.Get(tag(m.tag, mobile)).Result()
	return str != ""
}
// 删除
func (m *MobileCode) Del(mobile string) {
	m.redis.Del(tag(m.tag, mobile))
}