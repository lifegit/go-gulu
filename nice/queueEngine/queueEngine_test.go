/**
* @Author: TheLife
* @Date: 2021/5/29 下午3:27
 */
package queueEngine_test

import (
	"fmt"
	"github.com/lifegit/go-gulu/v2/nice/queueEngine"
	"testing"
	"time"
)

func TestQue(t *testing.T) {
	Que()

	<-time.After(time.Hour)
}

func BenchmarkQue(t *testing.B) {
	for range make([]int, t.N) {
		Que()
	}
}

func Que() {
	qu := queueEngine.NewEngine(1)

	qu.Customer(func(queue queueEngine.Queue) (b bool) {
		fmt.Println("receive start", queue)
		time.Sleep(time.Second)
		b = time.Now().UnixNano()%2 == 0
		fmt.Println("receive ok", queue, b)
		return
	})

	for key, _ := range make([]int, 3) {
		data := queueEngine.Queue{Data: fmt.Sprintf("第%d个", key+1)}
		qu.Producer(data)
		fmt.Println("send", data)
	}

	qu.ProducerFinish()
}
