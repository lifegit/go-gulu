/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package dbUtils

type DbWheres struct {
	Compare []Compare
	In      []In
	Range   []Range
	Like    []Like
}

// field = ?
// field > ? ||  field >= ?
// field < ? || field <= ?
const CompareEqual = "="
const CompareAboutEqual = ">="
const CompareAbout = ">"
const CompareLessEqual = "<="
const CompareLess = "<"

type Compare struct {
	Field string
	Type  string
	Text  interface{}
}

// field in(?)
// field not in(?)
type In struct {
	Not   bool
	Field string
	In    interface{}
}

// field >= start ANd field <= end
type Range struct {
	Field string
	Start int
	End   int
}

// field like %Text%
type Like struct {
	Field string
	Text  string
}
