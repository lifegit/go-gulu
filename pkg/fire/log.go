package fire

import (
	"context"
	"fmt"
	"gorm.io/gorm/logger"
	"time"
)

type Diary struct {
	Sql []string
}

func (n *Diary) LogMode(logger.LogLevel) logger.Interface {
	return &(*n)
}
func (n *Diary) Info(c context.Context, sql string, a ...interface{}) {
	fmt.Println("Info: ", sql)
}
func (n *Diary) Warn(c context.Context, sql string, a ...interface{}) {
	fmt.Println("Warn: ", sql)
}
func (n *Diary) Error(c context.Context, sql string, a ...interface{}) {
	fmt.Println("Error: ", sql)
}
func (n *Diary) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, _ := fc()
	n.Sql = append(n.Sql, sql)
	fmt.Println("Trace: ", sql)
}
func (n *Diary) LastSql(position ...int) string {
	p := If(position == nil, []int{1}, position).([]int)[0]
	if len(n.Sql) >= p {
		return n.Sql[len(n.Sql)-p]
	}

	return ""
}
