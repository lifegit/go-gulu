/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package person

import (
	"github.com/go-redis/redis"
	"time"
)

const DataSalt = "s"
const DataCode = "c"
const DataCaptcha = "v"
const DataIsCaptcha = "i"

// 用户数据
type Data struct {
	redis     *redis.Client
	personKey string
	isTtl     bool
	time      time.Duration
}

func NewData(personType string, personName string, personCode string, c *redis.Client) *Data {
	return &Data{
		personKey: "pd:" + personType + ":" + personName + ":" + personCode,
		isTtl:     false,
		time:      time.Second * 180,
		redis:     c,
	}
}

// 判断该用户是否存在
func (p *Data) IsEmpty() bool {
	if res, _ := p.redis.Exists(p.personKey).Result(); res == 1 {
		return true
	}
	return false
}

// 设置数据
func (p *Data) SetData(m map[string]interface{}) (bool, error) {
	if err := p.redis.HMSet(p.personKey, m).Err(); err != nil {
		return false, err
	}

	if !p.isTtl {
		if res, err := p.redis.PTTL(p.personKey).Result(); err == nil && res.Seconds() == -0.001 {
			p.redis.Expire(p.personKey, p.time)
		}
		p.isTtl = true
	}

	return true, nil
}

// 获取数据
func (p *Data) GetData(data []string) (map[string]interface{}, error) {
	res, err := p.redis.HMGet(p.personKey, data...).Result()
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	for i, val := range res {
		m[data[i]] = val
	}

	return m, err
}

// 删除用户
func (p *Data) Destroy() {
	p.redis.Del(p.personKey)
}
