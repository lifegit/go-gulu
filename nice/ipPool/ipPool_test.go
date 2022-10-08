package ipPool_test

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/lifegit/go-gulu/v2/nice/ipPool"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	ipMar := ipPool.New("serviceName", initRedis(), time.Hour*8,
		sourceFuncDemo,
	)

	ip, err := ipMar.Getter(true)
	assert.NoError(t, err)
	fmt.Println("one: ", ip, err)

	ipMar.Remove(ip)

	ip = ipMar.GetterWait(true)
	fmt.Println("two: ", ip, err)
}

func initRedis() (c *redis.Client) {
	c = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	if _, err := c.Ping().Result(); err != nil { // ok: pong == "PONG"
		log.Fatal(err)
	}

	return
}

// 获取ip来源demo
func sourceFuncDemo() (ips []string) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := httpClient.Get("")
	if err != nil {
		return
	}
	defer res.Body.Close()

	bytes, _ := ioutil.ReadAll(res.Body)
	str := string(bytes)

	return strings.Split(str, "\r\n")
}
