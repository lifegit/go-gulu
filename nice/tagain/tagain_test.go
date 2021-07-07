/**
* @Author: TheLife
* @Date: 2021/6/1 上午10:55
 */
package tagain

import (
	"fmt"
	"testing"
	"time"
)

func TestTAgain(t *testing.T) {
	first := true
	b := TAgain(func(n int) TryAgain {
		fmt.Println(n)
		if first {
			first = false
			return TryAgainFailNoTally
		} else if n <= 2 {
			return TryAgainFailTally
		}
		return TryAgainSuccess
	}, 3, time.Second)

	fmt.Println(b)
}
