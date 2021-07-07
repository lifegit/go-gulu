/**
* @Author: TheLife
* @Date: 2021/7/7 下午4:00
 */
package chromedpm_test

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/lifegit/go-gulu/v2/pkg/spider/chromedpm"
	"testing"
)

func TestChromeDpm(t *testing.T) {
	chromedpGetResponseBody()
}

func chromedpGetResponseBody() (err error) {
	ctx, cancel := chromedpm.NewAweChromeDp(10000, true)
	defer cancel()

	err = chromedp.Run(*ctx,
		chromedp.ActionFunc(func(cxx context.Context) error {
			chromedp.ListenTarget(*ctx, func(ev interface{}) {
				// 子请求完毕
				if ev, ok := ev.(*network.EventResponseReceived); ok {
					if ev.Response.Status == 200 {
						// 这里必须在协程内
						go func() {
							// 获取子请求的响应体
							if dataByte, err := network.GetResponseBody(ev.RequestID).Do(cxx); err != nil {
								dataStr := string(dataByte)
								fmt.Println(dataStr)
							}
						}()
					}

				}
			})
			return err
		}),
		chromedp.Navigate("https://www.baidu.com/"),
		chromedp.ActionFunc(func(ctx context.Context) (err error) {
			var s string
			err = chromedp.OuterHTML("body", &s).Do(ctx)
			return err
		}),
	)

	return
}

