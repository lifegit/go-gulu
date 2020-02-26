/**
* @Author: TheLife
* @Date: 2020-2-26 3:01 上午
 */
package cacheCount

import (
	"github.com/go-redis/redis"
	"time"
)

type Count struct {
	redis      *redis.Client
	expireTime time.Duration
	key        string
}

func New(key string, c *redis.Client, expireTime time.Duration) *Count {
	return &Count{
		key:        key,
		redis:      c,
		expireTime: expireTime,
	}
}
// get
func (p *Count) Get() (int, error) {
	return p.redis.Get(p.key).Int()
}
// add
func (p *Count) Add() (int64, error) {
	num, err := p.redis.Incr(p.key).Result()
	if err != nil {
		return 0, err
	}
	if num == 1 {
		p.redis.Expire(p.key, p.expireTime)
	}

	return num, nil
}
// set
func (p *Count) Set(s int) error {
	return p.redis.Set(p.key, s, p.expireTime).Err()
}
// del
func (p *Count) Destroy() {
	p.redis.Del(p.key)
}
