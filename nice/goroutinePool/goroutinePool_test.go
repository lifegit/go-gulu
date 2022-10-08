package goroutinePool_test

import (
	"fmt"
	"github.com/lifegit/go-gulu/v2/nice/goroutinePool"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	gors := goroutinePool.NewGoroutinePool(3)
	gors.Customer(func(tIndex int, data interface{}) {
		time.Sleep(time.Second * 3)
		fmt.Printf("index:%d, data:%v\n", tIndex, data)
	})
	gors.Producer(1)
	gors.Producer(2)
	gors.Wait()
	fmt.Println("@@")

	gors.Producer(3)
	gors.Producer(4)
	gors.Wait()

	fmt.Println("!!")
}
