/**
* @Author: TheLife
* @Date: 2020-11-8 6:23 下午
 */
package where

import "fmt"

// field in(?)
// field not in(?)
type In struct {
	Not   bool
	Field string
	In    interface{}
}

func (i *In) String() (query string, args interface{}) {
	not := ""
	if i.Not {
		not = "NOT"
	}
	return fmt.Sprintf("%s %s IN(?)", i.Field, not), i.In
}
