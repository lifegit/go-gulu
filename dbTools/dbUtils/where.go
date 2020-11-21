/**
* @Author: TheLife
* @Date: 2020-11-8 6:27 下午
 */
package dbUtils

import (
	"github.com/jinzhu/gorm"
	"go-gulu/dbTools/tool/where"
)

type DbWheres struct {
	Compare []where.Compare
	Range   []where.Range
	In      []where.In
	Like    []where.Like
}

func (d *DbUtils) WhereIsEmpty() bool {
	return d == nil || d.Wheres == nil ||
		( len(d.Wheres.Compare) == 0 && len(d.Wheres.Like) == 0 && len(d.Wheres.In) == 0 && len(d.Wheres.Range) == 0  )
}


func (d *DbUtils) Where(w interface{}) *DbUtils {
	if d.Wheres == nil {
		d.Wheres = &DbWheres{}
	}

	switch t := w.(type) {
	case where.Compare:
		d.Wheres.Compare = append(d.Wheres.Compare, t)
	case where.Range:
		d.Wheres.Range = append(d.Wheres.Range, t)
	case where.In:
		d.Wheres.In = append(d.Wheres.In, t)
	case where.Like:
		d.Wheres.Like = append(d.Wheres.Like, t)
	}

	return d
}

func (d *DbUtils) GetWhere(tx *gorm.DB) *gorm.DB {
	if d == nil || d.Wheres == nil {
		return tx
	}

	for _, value := range d.Wheres.Compare {
		tx = tx.Where(value.String())
	}
	for _, value := range d.Wheres.Range {
		tx = tx.Where(value.String())
	}
	for _, value := range d.Wheres.In {
		tx = tx.Where(value.String())
	}
	for _, value := range d.Wheres.Like {
		tx = tx.Where(value.String())
	}
	return tx
}
