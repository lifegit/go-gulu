/**
* @Author: TheLife
* @Date: 2020-3-19 4:03 上午
 */
package eventGroup

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	e := NewEventGroup()
	id := e.Register(time.Second * 3, 3 ,"hello", func(data interface{}) {
		fmt.Println(data)
	})

	e.Trigger(id)
	e.Trigger(id)
	//e.Trigger(id)

	time.Sleep(time.Second * 10)
}