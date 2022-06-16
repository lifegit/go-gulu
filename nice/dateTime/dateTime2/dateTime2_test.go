package dateTime2

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	now := time.Now()
	datetime := DateTime(now)

	// json format
	res1, _ := json.Marshal(datetime)
	res2, _ := json.Marshal(now.Format("2006-01-02 15:04:05"))
	assert.Equal(t, string(res1), string(res2))

	// call
	assert.Equal(t, now.Unix(), time.Time(datetime).Unix())
}
