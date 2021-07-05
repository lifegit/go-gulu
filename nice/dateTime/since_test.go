/**
* @Author: TheLife
* @Date: 2020-11-7 1:01 上午
 */
package dateTime

import (
	"testing"
	"time"
)

func TestName(t *testing.T) {
	var s Since
	s.Start()

	time.Sleep(time.Second)
	s.Dot(true)

	time.Sleep(time.Second)
	s.Dot(true)
}
