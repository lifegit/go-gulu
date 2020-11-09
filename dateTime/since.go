/**
* @Author: TheLife
* @Date: 2020-11-7 12:56 上午
 */
package dateTime

import (
	"fmt"
	"time"
)

type Since struct {
	Time time.Time
}

func (s *Since) Start() {
	s.Time = time.Now()
}
func (s *Since) Stop(isLog bool) {
	if sin := time.Since(s.Time); isLog {
		fmt.Println("since runtime:", sin)
	}
}
