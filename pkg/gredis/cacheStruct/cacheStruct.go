/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package cacheStruct

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"github.com/imdario/mergo"
	"reflect"
	"time"
)

type CacheStruct struct {
	key        string
	redis      *redis.Client
	expiration time.Duration
}

func New(key string, c *redis.Client, expiration ...time.Duration) *CacheStruct {
	e := time.Second * 180
	if expiration != nil {
		e = expiration[0]
	}
	return &CacheStruct{
		key:        key,
		redis:      c,
		expiration: e,
	}
}

// 判断该是否存在
func (p *CacheStruct) IsEmpty() bool {
	res, _ := p.redis.Exists(p.key).Result()

	return res == 1
}

// 第一次
func (p *CacheStruct) Init(m interface{}) (err error) {
	if p.IsEmpty() {
		return errors.New("empty")
	}
	return p.SetStruct(m)
}

// 设置数据
func (p *CacheStruct) SetStruct(m interface{}, merge ...bool) (err error) {
	// merge
	if merge != nil && merge[0] {
		v := reflect.New(reflect.ValueOf(m).Type()).Elem()
		if err = p.GetStruct(v); err == nil {
			if err = mergo.Merge(&m, v); err != nil {
				return
			}
		}
	}

	// set
	if err := p.redis.Set(p.key, m, p.expiration).Err(); err != nil {
		return err
	}

	return nil
}

// 获取数据
func (p *CacheStruct) GetStruct(callPoint interface{}) (err error) {
	res, err := p.redis.Get(p.key).Result()
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(res), callPoint)

	return err
}

// 删除用户
func (p *CacheStruct) Destroy() {
	p.redis.Del(p.key)
}
