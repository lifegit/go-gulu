package ipPool

import (
	"net/http"
	"net/url"
	"time"
)

// 验证
func (i *IpPool) verify(ip string) bool {
	// 1.验证cache
	if i.historyIsExist(ip) {
		return false
	}

	// 2. 验证可用性
	if !i.verifyAvailable(ip) {
		return false
	}

	return true
}

// 验证是否可连通
func (i *IpPool) verifyAvailable(proxyIp string) (b bool) {
	proxy, err := url.Parse(proxyIp)

	httpClient := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			Proxy:                 http.ProxyURL(proxy),
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * time.Duration(5),
		},
	}
	res, err := httpClient.Get("https://www.baidu.com")
	if err != nil {
		b = false
		//fmt.Println("错误信息：",err)
		return
	}
	defer res.Body.Close()

	return res.StatusCode == http.StatusOK
}
