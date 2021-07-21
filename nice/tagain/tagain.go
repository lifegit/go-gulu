/**
* @Author: TheLife
* @Date: 2021/6/1 上午10:54
 */
package tagain

import "time"

// 一个重试组件
type TryAgain int

const (
	TryAgainFailTally TryAgain = iota
	TryAgainFailNoTally
	TryAgainSuccess
)

func TAgain(fun func(n int) TryAgain, tryLen int, trySleep time.Duration) bool {
	for i := 0; i < tryLen; {
		switch fun(i) {
		case TryAgainSuccess:
			return true
		case TryAgainFailTally:
			i++
		case TryAgainFailNoTally:
		}

		time.Sleep(trySleep)
	}

	return false
}