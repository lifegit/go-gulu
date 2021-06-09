/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package order

import (
	"fmt"
)

type Type string

const (
	OrderAsc  Type = "asc"
	OrderDesc Type = "desc"
)

type Order struct {
	Field string `form:"name" binding:"required"`
	Type  Type   `binding:"required,eq=asc|eq=desc"`
}

func (o Order) String() string {
	return
}
