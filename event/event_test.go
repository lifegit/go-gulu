/**
* @Author: TheLife
* @Date: 2020-3-19 3:07 上午
 */
package event

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	event := NewEvent()
	id := event.Register(time.Second * 5,"hello", func(data interface{}) {
		fmt.Println(data)
	})

	time.Sleep(time.Second * 3)
	event.Trigger(id)


	time.Sleep(time.Minute * 10)
}

