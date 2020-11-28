/**
* @Author: TheLife
* @Date: 2020/11/28 下午2:10
 */
package rwMutexLog

import (
	"testing"
	"time"
)

func TestRWMutex(t *testing.T) {
	var m RWMutex

	go m.Unlock()
	m.Lock()

	time.Sleep(time.Second * 3)
}
