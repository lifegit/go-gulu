/**
* @Author: TheLife
* @Date: 2020-11-8 6:24 下午
 */
package update

import (
	"fmt"
	"gorm.io/gorm"
)

// field = field Type Number # field = field + 1
type ArithmeticType string

const (
	ArithmeticIncrease ArithmeticType = "+"
	ArithmeticReduce   ArithmeticType = "-"
	ArithmeticMultiply ArithmeticType = "*"
	ArithmeticExcept   ArithmeticType = "/"
)

type Arithmetic struct {
	Field  string
	Type   ArithmeticType
	Number float32
}

func (a *Arithmetic) String() (val interface{}) {
	return gorm.Expr(fmt.Sprintf("`%s` %s ?", a.Field, a.Type), a.Number)
}
