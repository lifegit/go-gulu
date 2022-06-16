package queueEngine_test

import (
	"fmt"
	"github.com/lifegit/go-gulu/v2/nice/queueEngine"
	"testing"
	"time"
)

func TestQue(t *testing.T) {
	Que()
}

func BenchmarkQue(t *testing.B) {
	for range make([]int, t.N) {
		Que()
	}
}

func Que() {
	qu := queueEngine.NewEngine(2)

	go func() {
		for key, _ := range make([]int, 10) {
			data := queueEngine.Queue{Data: fmt.Sprintf("第%d个", key+1)}
			qu.Producer(data)
			fmt.Println("send", data)
		}

		qu.ProducerFinish()
	}()

	qu.CustomerWait(func(queue queueEngine.Queue) (b bool) {
		time.Sleep(time.Second)
		b = time.Now().UnixNano()%2 == 0
		fmt.Println("receive ok", queue, b)
		return
	})
}
