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

func (s *Since) Dot(isLog bool) {
	if sin := time.Since(s.Time); isLog {
		fmt.Println("since runtime:", sin)
	}
}
