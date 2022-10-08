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
