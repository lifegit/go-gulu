/**
* @Author: TheLife
* @Date: 2020/11/28 下午2:10
 */
package rwMutexLog_test

import (
	"fmt"
	"github.com/lifegit/go-gulu/sync/rwMutexLog"
	"testing"
)

func TestRWMutex(t *testing.T) {
	var m rwMutexLog.RWMutex

	func() {
		m.Lock()
		defer m.Unlock()

		fmt.Println("hello")
	}()
}
