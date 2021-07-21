/**
* @Author: TheLife
* @Date: 2021/5/18 下午2:26
 */
package koa_test

import (
	"fmt"
	"github.com/lifegit/go-gulu/v2/nice/koa"
	"github.com/lifegit/go-gulu/v2/nice/koa/koaMiddleware"
	"testing"
	"time"
)

func TestName2(t *testing.T) {
	c := koa.NewContext()
	c.Use(koaMiddleware.Recovery(), workFunc, koaMiddleware.NewLoggerMiddlewareSmoothFail(true, true, "testKoa", "./log"))

	go func() {
		for key := range make([]int, 5) {
			fmt.Println(key)
			go c.Run()
			time.Sleep(time.Millisecond * 1500)
		}
	}()

	time.Sleep(time.Minute)
}

func workFunc(c *koa.Context) {
	fmt.Println("working...")

	if time.Now().Unix()%2 != 0 {
		panic("1")
	} else {
		fmt.Println("work success!")
		c.Result.Data = time.Now().Unix()
		c.Next()

		time.Sleep(time.Second * 3)
		fmt.Println("work distracted!")
	}
}
