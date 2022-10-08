package dateTime

import "time"

type UnixRangeType uint8

const (
	Month UnixRangeType = iota
	Day
)

// 获取一个时间的开始与结束的时间戳
func UnixRange(timeType UnixRangeType, t time.Time) (start, end int64) {
	year, month, day := t.Date()
	if timeType == Month { // 本月开始与结束
		thisMonthFirst := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
		thisMonthLast := thisMonthFirst.AddDate(0, 1, -1)
		start = thisMonthFirst.Unix()
		end = thisMonthLast.Unix() + 86400 - 1
	} else if timeType == Day { // 本天开始与结束
		thisDayFirst := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
		start = thisDayFirst.Unix()
		end = thisDayFirst.Unix() + 86400 - 1
	}
	return
}
