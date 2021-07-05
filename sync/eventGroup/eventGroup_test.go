/**
* @Author: TheLife
* @Date: 2020-3-19 4:03 上午
 */
package eventGroup_test

import (
	"fmt"
	"github.com/lifegit/go-gulu/v2/sync/eventGroup"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	ch := make(chan bool)

	e := eventGroup.NewEventGroup()
	id := e.Register(time.Second*3, 3, "hello", func(data interface{}) {
		fmt.Println(data)
		ch <- true
	})

	e.Trigger(id)
	e.Trigger(id)
	e.Trigger(id)

	<-ch
}
