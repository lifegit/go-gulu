/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package person

import (
	"github.com/go-redis/redis"
	"github.com/lifegit/go-gulu/pkg/gredis/cacheCount"
	"time"
)

// Ip
type Ip struct {
	cache *cacheCount.Count
	max   int
}

func NewIp(ip string, tryMax int, expireTimeMinute int, c *redis.Client) *Ip {
	return &Ip{
		cache: cacheCount.New("pi:"+ip, c, time.Minute*time.Duration(expireTimeMinute)),
		max:   tryMax,
	}
}

// 是否频繁
func (c *Ip) IsBusy() bool {
	res, err := c.cache.Get()
	if err != nil {
		return false
	}
	return res >= c.max
}

// 增加一个次数,并且返回增加次数后是否频繁
func (c *Ip) AddCount() (bool, error) {
	num, err := c.cache.Add()

	return num >= int64(c.max), err
}

// 删除
func (c *Ip) Destroy() {
	c.cache.Destroy()
}
