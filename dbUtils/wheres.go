/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package dbUtils

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

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

func (utils DbUtils) GetWhere(tx *gorm.DB) *gorm.DB {
	where := utils.Where
	if where != nil {
		for _, value := range where.Compare {
			tx = tx.Where(fmt.Sprintf("%s %s ?", value.Field, value.Type), value.Text)
		}
		for _, value := range where.In {
			n := ""
			if value.Not {
				n = "not"
			}
			tx = tx.Where(fmt.Sprintf("%s %s in(?)", value.Field, n), value.In)
		}

		for _, value := range where.Range {
			tx = tx.Where(fmt.Sprintf("%s >= ? AND %s <= ?", value.Field, value.Field), value.Start, value.End)
		}

		for _, value := range where.Like {
			tx = tx.Where(value.Field+" LIKE ?", fmt.Sprintf("%%%s%%", value.Text))
		}
	}
	return tx
}
