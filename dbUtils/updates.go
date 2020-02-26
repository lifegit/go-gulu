/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package dbUtils

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
