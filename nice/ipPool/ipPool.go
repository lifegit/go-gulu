package ipPool

import (
	"errors"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

// ip 池

type IpPool struct {
	tag    string
	redis  *redis.Client
	expire time.Duration
	mtx    sync.RWMutex

	source Source
}

func New(tag string, redis *redis.Client, expire time.Duration, sourceFunc ...Func) *IpPool {
	i := &IpPool{
		tag:    tag,
		redis:  redis,
		expire: expire,
	}
	i.source.AddSource(sourceFunc...)

	return i
}

// 获取IP
func (i *IpPool) Getter(wayRemove bool) (ip string, err error) {
	i.mtx.Lock()
	defer i.mtx.Unlock()

	defer func() {
		if ip != "" && wayRemove {
			i.Remove(ip)
		}
	}()

	popSlicesAndVerify := func() (ip string, err error) {
		for true {
			ip, err = i.slicesGet()
			if ip, err = i.slicesGet(); err != nil || i.verify(ip) {
				return
			}
		}
		return
	}

	if ip, err = popSlicesAndVerify(); ip != "" {
		return
	}
	for _, item := range i.source.GetFuncResult() {
		if res := item(); len(res) > 0 {
			i.slicesPush(res...)
			if ip, err = popSlicesAndVerify(); ip != "" {
				return
			}
		}
	}

	if ip == "" {
		err = errors.New("没有可用IP")
	}

	return
}

// 获取ip，没有则阻塞，直到获取到
func (i *IpPool) GetterWait(wayRemove bool) (ip string) {
	for true {
		if ip, _ = i.Getter(wayRemove); ip != "" {
			return
		}

		time.Sleep(time.Second * 30)
	}

	return
}

// 释放
func (i *IpPool) Remove(ip string) {
	if i.slicesRemove(ip) {
		i.historyPush(ip)
	}
}
