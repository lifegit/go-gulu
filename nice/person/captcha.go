/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package person

import (
	"github.com/go-redis/redis"
	"github.com/lifegit/go-gulu/v2/pkg/gredis/cacheCount"
	"time"
)

// 验证码
type Captcha struct {
	cache *cacheCount.Count
	max   int
}

// 类型，名称，redis，最大数量，过期时间分钟
func NewCaptcha(personType string, personName string, tryMax int, expireTimeMinute int, c *redis.Client) *Captcha {
	return &Captcha{
		cache: cacheCount.New("pc:"+personType+":"+personName, c, time.Minute*time.Duration(expireTimeMinute)),
		max:   tryMax,
	}
}

// 是否频繁
func (c *Captcha) IsBusy() bool {
	res, err := c.cache.Get()
	if err != nil {
		return false
	}
	return res >= c.max
}

// 增加一个次数,并且返回增加次数后是否频繁
func (c *Captcha) AddCount() (bool, error) {
	num, err := c.cache.Add()

	return num >= int64(c.max), err
}

// 删除
func (c *Captcha) Destroy() {
	c.cache.Destroy()
}
