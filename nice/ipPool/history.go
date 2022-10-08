package ipPool

import (
	"fmt"
)

// 是否存在
func (i *IpPool) historyIsExist(ip string) bool {
	res, _ := i.redis.Exists(i.historyKey(ip)).Result()

	return res == 1
}

// 添加
func (i *IpPool) historyPush(ip string) bool {
	err := i.redis.Set(i.historyKey(ip), "", i.expire).Err()

	return err == nil
}

// 构造key
func (i *IpPool) historyKey(ip string) string {
	return fmt.Sprintf("ip:%s:hs:%s", i.tag, ip)
}
