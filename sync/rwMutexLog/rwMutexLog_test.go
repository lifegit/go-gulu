package rwMutexLog_test

import (
	"fmt"
	"github.com/lifegit/go-gulu/v2/sync/rwMutexLog"
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
