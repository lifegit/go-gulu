package dateTime2

import (
	"fmt"
	"time"
)

// 每次调用都要先转回time.Time，所以不推荐
// https://www.cnblogs.com/xiaofengshuyu/p/5664654.html

type DateTime time.Time

func (t *DateTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	*t = DateTime(now)
	return
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t DateTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}
