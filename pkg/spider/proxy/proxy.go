/**
* @Author: TheLife
* @Date: 2021/6/1 上午10:04
 */
package proxy

import "fmt"

type Proxy string
type Schema string

const (
	SchemaHttp           = "http://"
	SchemaSocket5 Schema = "socks5://"
)

func NewProxy(s Schema, userAndPass string, addrPort string) Proxy {
	// 规则：
	// https://github.com/chromedp/chromedp/issues/190
	// Schema[<user>:<password>@]<host>:<port>
	// 示例：
	// socks5://192.168.1.1:8080
	// socks5://username:password@192.168.1.1:8080
	// socks5://username:password@proxyserver.com:31280

	// http://192.168.1.1:8080
	// http://username:password@192.168.1.1:8080
	// http://username:password@proxyserver.com:31280

	if userAndPass != "" {
		userAndPass += "@"
	}

	return Proxy(fmt.Sprintf("%s%s%s", s, userAndPass, addrPort))
}

func (p Proxy) ToString() string {
	return string(p)
}