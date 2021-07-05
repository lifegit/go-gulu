package dateTime1

import (
	"fmt"
	"time"
)

// 推荐
// http://www.axiaoxin.com/article/241/
type DateTime struct {
	time.Time
}

func (t *DateTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	*t = DateTime{now}
	return
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t DateTime) String() string {
	return t.Format("2006-01-02 15:04:05")
}
