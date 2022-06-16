package dateTime

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRange(t *testing.T) {
	u := time.Unix(1623902400, 0) // 2021-06-17 12:00:00

	// month
	start, end := UnixRange(Month, u)
	fmt.Println(
		time.Unix(start, 0).String(),
		time.Unix(end, 0).String(),
	)
	assert.Equal(t, start, int64(1622476800)) // 2021-06-01 00:00:00
	assert.Equal(t, end, int64(1625068799))   // 2021-06-30 23:59:59

	// day
	start, end = UnixRange(Day, u)
	fmt.Println(
		time.Unix(start, 0).String(),
		time.Unix(end, 0).String(),
	)
	assert.Equal(t, start, int64(1623859200)) // 2021-06-17 00:00:00
	assert.Equal(t, end, int64(1623945599))   // 2021-06-17 23:59:59
}
