// Package fire /**
// 常见的curd工具集，快速助力业务开发。

package fire

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
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

type validateOnce struct {
	*validator.Validate
	once sync.Once
}

var validate validateOnce

// CrudCreate
// model support Array,Slice,Struct
// every struct support tag `gormCreate` validator.Struct()
func (d *Fire) CrudCreate(model interface{}, batchSize ...int) (err error) {
	validate.once.Do(func() {
		validate.Validate = validator.New()
		validate.SetTagName("gormCreate")
	})

	reflectValue := reflect.Indirect(reflect.ValueOf(model))
	switch reflectValue.Kind() {
	case reflect.Array, reflect.Slice:
		// check
		for i := 0; i < reflectValue.Len(); i++ {
			o := reflectValue.Index(i).Interface()
			if err = validate.Struct(o); err != nil {
				return
			}
		}
		// batchSize
		batch := 20
		if batchSize != nil {
			batch = batchSize[0]
		} else if d.CreateBatchSize != 0 {
			batch = d.CreateBatchSize
		}
		d.CreateInBatches(model, batch)
	default:
		// check
		if err = validate.Struct(model); err != nil {
			return
		}
		err = d.Create(model).Error
	}

	return
}

func (d *Fire) IsExists(model interface{}) bool {
	v := reflect.New(reflect.ValueOf(model).Type()).Elem()
	tx := d.Model(model).Where(model).Take(&v)

	return tx.RowsAffected >= 1
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
	defer pageResult.SetData(callListData)

	d.Model(model).Where(model).Count(&pageResult.Total)
	if pageResult.Total > 0 || d.DryRun {
		d.Statement.SQL = strings.Builder{}
		d.Offset(pageResult.GetOffset()).Limit(pageResult.PageSize).Find(callListData)
	}

	return
}

func (d *Fire) CrudOnePreloadJoin(model interface{}, callData interface{}) (err error) {
	return d.PreloadJoin(model).Where(model).Take(callData).Error
}

func (d *Fire) CrudAllPreloadJoin(model interface{}, callListData interface{}) (err error) {
	return d.PreloadJoin(model).Where(model).Find(callListData).Error
}

func (d *Fire) CrudAllPagePreloadJoin(model interface{}, callListData interface{}, page ...Page) (pageResult PageResult, err error) {
	pageResult.Init(page...)
	defer pageResult.SetData(callListData)

	tx := d.PreloadJoin(model).Model(model).Where(model)
	tx.Session(&gorm.Session{}).Count(&pageResult.Total)
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
	defer pageResult.SetData(callListData)

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
	err = d.Model(model).Where(model).Select(d.SumLabel(column)).Find(&sum).Error

	return
}

type M map[string]interface{}

// CrudUpdate
// updates support (M or map[string]interface{}) and struct
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
					v, _ := field.ValueOf(d.Statement.Context, reflect.ValueOf(item))
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

// CrudUpdateByPrimaryKey Make sure that all primary keys are not zero when updating
func (d *Fire) CrudUpdateByPrimaryKey(model interface{}, updates ...interface{}) (err error) {
	sch, err := schema.Parse(model, &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		return
	}

	for _, field := range sch.PrimaryFields {
		if _, isZero := field.ValueOf(d.Statement.Context, reflect.ValueOf(model)); isZero {
			return errors.New(fmt.Sprintf("primary key %s is not exist", field.Name))
		}
	}

	tx := NewInstance(d.Limit(1))
	return tx.CrudUpdate(model, updates...)
}

func (d *Fire) CrudDelete(model interface{}) (err error) {
	// WARNING: When there is no condition WHERE, error is gorm.ErrMissingWhereClause("WHERE conditions required"), so you can safely delete it.
	err = d.Where(model).Delete(model).Error

	return
}
