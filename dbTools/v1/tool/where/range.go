/**
* @Author: TheLife
* @Date: 2020-11-8 6:23 下午
 */
package where

import "fmt"

// field >= start ANd field <= end
type Range struct {
	Field string
	Start int64
	End   int64
}

func (r *Range) String() (query string, start, end int64) {
	return fmt.Sprintf("`%s` >= ? AND `%s` <= ?", r.Field, r.Field), r.Start, r.End
}
