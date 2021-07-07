/**
* @Author: TheLife
* @Date: 2021/7/7 下午3:19
 */
package proxy_test

import (
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/lifegit/go-gulu/v2/pkg/spider/chromedpm"
	"github.com/lifegit/go-gulu/v2/pkg/spider/proxy"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProxy(t *testing.T) {
	p := proxy.NewProxy(proxy.SchemaSocket5, "", "192.168.1.1:8080")
	assert.Equal(t, p.ToString(), "socks5://192.168.1.1:8080")

	p = proxy.NewProxy(proxy.SchemaSocket5, "username:password", "192.168.1.1:8080")
	assert.Equal(t, p.ToString(), "socks5://username:password@192.168.1.1:8080")

	p = proxy.NewProxy(proxy.SchemaSocket5, "username:password", "proxyserver.com:31280")
	assert.Equal(t, p.ToString(), "socks5://username:password@proxyserver.com:31280")


	p = proxy.NewProxy(proxy.SchemaHttp, "", "192.168.1.1:8080")
	assert.Equal(t, p.ToString(), "http://192.168.1.1:8080")

	p = proxy.NewProxy(proxy.SchemaHttp, "username:password", "192.168.1.1:8080")
	assert.Equal(t, p.ToString(), "http://username:password@192.168.1.1:8080")

	p = proxy.NewProxy(proxy.SchemaHttp, "username:password", "proxyserver.com:31280")
	assert.Equal(t, p.ToString(), "http://username:password@proxyserver.com:31280")
}

func TestTool(t *testing.T) {
	_ = collyProxy(proxy.NewProxy(proxy.SchemaSocket5, "", "192.168.1.1:8080"))
	_ = chromeDpProxy()
}


func collyProxy (proxy ...proxy.Proxy) (err error) {
	c := colly.NewCollector()
	if proxy != nil {
		_ = c.SetProxy(proxy[0].ToString())
	}

	c.OnHTML("body", func(element *colly.HTMLElement) {

	})

	if e := c.Visit("https://www.baidu.com/"); e != nil{
		err = e
	}

	return
}

func chromeDpProxy(proxy ...proxy.Proxy) (err error) {
	var c []chromedp.ExecAllocatorOption
	if proxy != nil {
		c = append(c, chromedp.ProxyServer(proxy[0].ToString()))
	}

	ctx, cancel := chromedpm.NewAweChromeDp(10000, true, c...)
	defer cancel()

	err = chromedp.Run(*ctx,
		chromedp.Navigate("https://www.baidu.com/"),
	)

	return
}