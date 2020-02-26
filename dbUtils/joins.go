/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package dbUtils

type JoinOn struct {
	Table string
	Field string
}
type LeftJoin struct {
	TableName string
	Left      JoinOn
	Right     JoinOn
}
type DbJoin struct {
	LeftJoin []LeftJoin
}
