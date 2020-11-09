/**
* @Author: TheLife
* @Date: 2020-11-8 6:24 下午
 */
package where

import "fmt"

// field like %Text%
type Like struct {
	Field string
	Text  string
}

func (l *Like) String() (query string, args string) {
	return fmt.Sprintf("`%s` LIKE ?", l.Field), fmt.Sprintf("%%%s%%", l.Text)
}
