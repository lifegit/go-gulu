package ipPool

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

// 添加
func (i *IpPool) slicesPush(ip ...string) bool {
	// 过滤非法Ip
	res := make([]string, 0)
	for _, value := range ip {
		if address := net.ParseIP(strings.Split(value, ":")[0]); address != nil {
			res = append(res, value)
		}
	}

	err := i.redis.RPush(i.slicesKey(), res).Err()
	if err == nil {
		i.redis.Expire(i.slicesKey(), i.expire)
	}

	return err == nil
}

// 获取
func (i *IpPool) slicesGet() (ip string, err error) {
	res := i.slicesKey()

	ip, err = i.redis.LIndex(res, 0).Result()
	if ip == "" && err == nil {
		err = errors.New("不存在ip")
	}

	return
}

// 删除
func (i *IpPool) slicesRemove(ip string) bool {
	err := i.redis.LRem(i.slicesKey(), 1, ip).Err()

	return err == nil
}

func (i *IpPool) slicesKey() string {
	return fmt.Sprintf("ip:%s:sc", i.tag)
}
