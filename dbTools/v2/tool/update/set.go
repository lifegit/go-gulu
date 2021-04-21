/**
* @Author: TheLife
* @Date: 2020-11-8 6:25 下午
 */
package update

type Set struct {
	Field string
	Value interface{}
}

func (a *Set) String() (val interface{}) {
	return a.Value
}
