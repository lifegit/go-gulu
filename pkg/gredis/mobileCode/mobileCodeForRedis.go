/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package mobileCode

import (
	"github.com/go-redis/redis"
	"time"
)

// sms := mobileCode.New(mobileCode.ForRedis{ Store: app.Cache, Expire: time.Minute * 5, TagPrefix: "m:cp:" } )

type ForRedis struct {
	TagPrefix string
	Expire    time.Duration
	Store     *redis.Client
}

func (f ForRedis) Set(key, value string) error {
	return f.Store.Set(tag(f.TagPrefix, key), value, f.Expire).Err()
}

func (f ForRedis) Get(key string) (string, error) {
	return f.Store.Get(tag(f.TagPrefix, key)).Result()
}

func (f ForRedis) Del(key string) error {
	return f.Store.Del(tag(f.TagPrefix, key)).Err()
}
