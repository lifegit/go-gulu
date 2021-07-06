/**
* @Author: TheLife
* @Date: 2021/5/27 下午5:08
 */
package fire

import (
	"errors"
	"fmt"
	"github.com/imdario/mergo"
	"github.com/lifegit/go-gulu/v2/conv/arrayconv"
	"github.com/lifegit/go-gulu/v2/conv/structureconv"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"reflect"
	"strings"
	"sync"
)

// 常见的curd工具集，快速助力业务开发。

func (d *Fire) IsExists(model interface{}) bool {
	var s interface{}
	d.Model(model).Where(model).Select("1").Take(&s)

	return s != nil
}

func (d *Fire) CrudOne(model interface{}, callData interface{}) (err error) {
	err = d.Model(model).Where(model).Take(callData).Error

	return
}

func (d *Fire) CrudAll(model interface{}, callListData interface{}) (err error) {
	err = d.Model(model).Where(model).Find(callListData).Error

	return
}

func (d *Fire) CrudAllPage(model interface{}, callListData interface{}, page ...Page) (pageResult PageResult, err error) {
	pageResult.Init(page...)

	d.Model(model).Where(model).Count(&pageResult.Total)
	if pageResult.Total > 0 || d.DryRun {
		d.Statement.SQL = strings.Builder{}
		d.Offset(pageResult.GetOffset()).Limit(pageResult.PageSize).Find(callListData)
	}

	return
}

func (d *Fire) CrudOnePreloadJoin(model interface{}, callData interface{}) (err error) {
	return d.PreloadJoin(model).Take(callData).Error
}

func (d *Fire) CrudAllPreloadJoin(model interface{}, callListData interface{}) (err error) {
	return d.PreloadJoin(model).Find(callListData).Error
}

func (d *Fire) CrudAllPagePreloadJoin(model interface{}, callListData interface{}, page ...Page) (pageResult PageResult, err error) {
	pageResult.Init(page...)

	tx := d.PreloadJoin(model).Session(&gorm.Session{})
	tx.Model(model).Count(&pageResult.Total)
	if pageResult.Total > 0 || d.DryRun {
		tx.Offset(pageResult.GetOffset()).Limit(pageResult.PageSize).Find(callListData)
	}

	return
}

func (d *Fire) CrudOnePreloadAll(model interface{}, callData interface{}) (err error) {
	return d.PreloadAll().Where(model).Take(callData).Error
}

func (d *Fire) CrudAllPreloadAll(model interface{}, callListData interface{}) (err error) {
	return d.PreloadAll().Where(model).Find(callListData).Error
}

func (d *Fire) CrudAllPagePreloadAll(model interface{}, callListData interface{}, page ...Page) (pageResult PageResult, err error) {
	pageResult.Init(page...)

	tx := d.PreloadAll().Model(model).Where(model)
	tx.Count(&pageResult.Total)
	if pageResult.Total > 0 || d.DryRun {
		tx.Offset(pageResult.GetOffset()).Limit(pageResult.PageSize).Find(callListData)
	}

	return
}

func (d *Fire) CrudCount(model interface{}) (count int64, err error) {
	err = d.Model(model).Where(model).Count(&count).Error

	return
}

func (d *Fire) CrudSum(model interface{}, column string) (sum float32, err error) {
	err = d.Model(model).Where(model).Select(fmt.Sprintf("IFNULL(SUM(`%s`),0)", column)).Find(&sum).Error

	return
}

type M map[string]interface{}

// updates support M(or map[string]interface{}) and struct
// support gorm.Db.Select() and gorm.Db.Omit()
// TODO: struct only update non-zero fields
func (d *Fire) CrudUpdate(model interface{}, updates ...interface{}) (err error) {
	// toMap
	m := make(M)
	mTemp := make(M)
	for _, item := range updates {
		switch reflect.TypeOf(item).Kind() {
		case reflect.Struct:
			if s, err := schema.Parse(item, &sync.Map{}, schema.NamingStrategy{}); err == nil {
				for _, field := range s.Fields {
					v, _ := field.ValueOf(reflect.ValueOf(item))
					m[field.DBName] = v
				}
			}
		case reflect.Map:
			var mMap M
			switch item.(type) {
			case M:
				mMap = item.(M)
			case map[string]interface{}:
				mMap = item.(map[string]interface{})
			}

			if err = mergo.Merge(&mTemp, mMap); err != nil {
				return
			}
			if err = mergo.Merge(&m, mMap); err != nil {
				return
			}
		default:
			return gorm.ErrInvalidValue
		}
	}
	// omits
	for _, omitItem := range d.Statement.Omits {
		if m[omitItem] != nil {
			delete(m, omitItem)
		}
	}

	if len(d.Statement.Selects) > 0 {
		// select
		if !arrayconv.StringIn("*", d.Statement.Selects) {
			// delete if not found
			for key, _ := range m {
				if !arrayconv.StringIn(key, d.Statement.Selects) {
					delete(m, key)
				}
			}
		}
	} else {
		// struct isBlank
		for key, item := range m {
			if mTemp[key] == nil && structureconv.IsBlank(item) {
				delete(m, key)
			}
		}
	}
	if len(m) <= 0 {
		return errors.New("updates conditions required")
	}

	tx := d.Model(model).Where(model)
	if e, ok := tx.Statement.Clauses["WHERE"].Expression.(clause.Where); !ok || len(e.Exprs) <= 0 {
		return gorm.ErrMissingWhereClause
	}

	tx = tx.Updates(map[string]interface{}(m))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected <= 0 {
		return errors.New("resource is not found")
	}

	return tx.Error
}

func (d *Fire) CrudDelete(model interface{}) (err error) {
	// WARNING: When there is no condition WHERE, error is gorm.ErrMissingWhereClause("WHERE conditions required"), so you can safely delete it.
	err = d.Where(model).Delete(model).Error

	return
}
