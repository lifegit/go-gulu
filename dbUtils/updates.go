/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package dbUtils

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type DbUpdates struct {
	Arithmetic []Arithmetic
	Set        []Set
}

// field = field Type Number // field = field + 1
const ArithmeticIncrease = "+"
const ArithmeticReduce = "-"
const ArithmeticMultiply = "*"
const ArithmeticExcept = "/"

type Arithmetic struct {
	Field  string
	Type   string
	Number float32
}

type Set struct {
	Field string
	Value interface{}
}

func (utils DbUtils) GetUpdate() (m *map[string]interface{}) {
	updates := utils.Updates
	if updates != nil {
		maps := make(map[string]interface{})
		for _, value := range updates.Arithmetic {
			maps[value.Field] = gorm.Expr(fmt.Sprintf("%s %s ?", value.Field, value.Type), value.Number)
		}
		for _, value := range updates.Set {
			maps[value.Field] = value.Value
		}
		return &maps
	}
	return nil
}
