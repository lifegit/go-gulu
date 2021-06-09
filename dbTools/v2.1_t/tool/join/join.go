/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package join

import (
	"fmt"
)

type JoinOn struct {
	Table string
	Field string
}

func (j JoinOn) String() string {
	if j.Table != "" {
		return fmt.Sprintf("`%s`.%s", j.Table, j.Field)
	} else {
		return j.Field
	}
}

type LeftJoin struct {
	Left  JoinOn
	Right JoinOn
}

func (l *LeftJoin) String() (val string) {
	return fmt.Sprintf("LEFT JOIN %s ON %s = %s", l.Right.Table, l.Left, l.Right)
}
