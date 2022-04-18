// Package fire /**
// 对 gorm.DB 的补充封装，常见的curd工具集，快速助力业务开发。实现更爽快得使用。属于基础层服务代码。

package fire

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/imdario/mergo"
	"github.com/lifegit/go-gulu/v2/conv/arrayconv"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"net/url"
	"reflect"
	"sync"
)

type Fire struct {
	*gorm.DB
}

func NewInstance(db *gorm.DB) *Fire {
	return &Fire{DB: db}
}

// Close DB
func (d *Fire) Close() (err error) {
	dbs, err := d.DB.DB()
	if err != nil {
		return
	}

	return dbs.Close()
}

// ModelWhere
func (d *Fire) ModelWhere(model interface{}) *Fire {
	tx := d.Model(model).Where(model)

	return NewInstance(tx)
}

// PreloadJoin
// TODO：Single SQL, mysql bonding data, so the conditions of all query tables are supported. use Join you need to pay attention to performance
func (d *Fire) PreloadJoin(model interface{}) *Fire {
	if reflect.TypeOf(model).Kind() != reflect.Struct {
		return nil
	}
	tx := d.DB
	key := reflect.TypeOf(model)
	val := reflect.ValueOf(model)
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Kind() == reflect.Struct {
			if val.Field(i).CanInterface() {
				// anonymous not join
				if !key.Field(i).Anonymous {
					tx = tx.Joins(key.Field(i).Name)
				}
			}
		}
	}

	return NewInstance(tx)
}

// PreloadAll
// TODO：Multiple SQL, gorm bonding data, so query conditions other than the main table are not supported
func (d *Fire) PreloadAll() *Fire {
	tx := d.DB.Preload(clause.Associations)

	return NewInstance(tx)
}

func (d *Fire) Allow(values url.Values, allow Allow) *Fire {
	tx := d.DB
	tx = allow.AllowSort(values, tx)
	tx = allow.AllowParams(values, tx)
	return NewInstance(tx)
}

func (d *Fire) Clause(exps ...clause.Expression) *Fire {
	return NewInstance(d.Clauses(exps...))
}

func (d *Fire) OrderByColumn(column interface{}, order OrderType) *Fire {
	return NewInstance(d.Order(Order(column, order)))
}

type validateOnce struct {
	*validator.Validate
	once sync.Once
}

var validate validateOnce

func (d *Fire) IsExists(model interface{}) bool {
	v := reflect.New(reflect.ValueOf(model).Type()).Elem()
	tx := d.Model(model).Where(model).Take(&v)

	return tx.RowsAffected >= 1
}

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

func (d *Fire) CrudOne(model interface{}, callData interface{}) (err error) {
	err = d.Model(model).Where(model).Take(callData).Error

	return
}

func (d *Fire) CrudAll(model interface{}, callListData interface{}) (err error) {
	err = d.Model(model).Where(model).Find(callListData).Error

	return
}

func (d *Fire) crudAllPage(model interface{}, callListData interface{}, page interface{}, countBeforeFunc func(*gorm.DB) *gorm.DB, findBeforeFunc func(*gorm.DB) *gorm.DB) (pageResult PageResult, err error) {
	pageResult.Init(page)
	defer pageResult.SetData(callListData)

	tx := d.Model(model).Where(model)
	if countBeforeFunc != nil {
		tx = countBeforeFunc(tx)
	}
	err = tx.Session(&gorm.Session{}).Select([]string{}).Count(&(pageResult.Total)).Error
	if err != nil {
		return
	}
	if pageResult.Total > 0 || tx.DryRun {
		if findBeforeFunc != nil {
			tx = findBeforeFunc(tx)
		}
		err = tx.Offset(pageResult.GetOffset()).Limit(pageResult.PageSize).Find(callListData).Error
	}

	return
}

func (d *Fire) CrudAllPage(model interface{}, callListData interface{}, page interface{}) (pageResult PageResult, err error) {
	return d.crudAllPage(model, callListData, page, nil, nil)
}

func (d *Fire) CrudOnePreloadJoin(model interface{}, callData interface{}) (err error) {
	return d.PreloadJoin(model).Where(model).Take(callData).Error
}

func (d *Fire) CrudAllPreloadJoin(model interface{}, callListData interface{}) (err error) {
	return d.PreloadJoin(model).Where(model).Find(callListData).Error
}

func (d *Fire) CrudAllPagePreloadJoin(model interface{}, callListData interface{}, page interface{}) (pageResult PageResult, err error) {
	var preloads map[string][]interface{}
	return d.crudAllPage(model, callListData, page, func(db *gorm.DB) *gorm.DB {
		tx := NewInstance(db).PreloadJoin(model).DB
		preloads = tx.Statement.Preloads
		tx.Statement.Preloads = map[string][]interface{}{}
		return tx
	}, func(db *gorm.DB) *gorm.DB {
		if len(preloads) > 0 {
			db.Statement.Preloads = preloads
		}
		return db
	})
}

func (d *Fire) CrudOnePreloadAll(model interface{}, callData interface{}) (err error) {
	return d.PreloadAll().Where(model).Take(callData).Error
}

func (d *Fire) CrudAllPreloadAll(model interface{}, callListData interface{}) (err error) {
	return d.PreloadAll().Where(model).Find(callListData).Error
}

func (d *Fire) CrudAllPagePreloadAll(model interface{}, callListData interface{}, page interface{}) (pageResult PageResult, err error) {
	return d.crudAllPage(model, callListData, page, nil, func(db *gorm.DB) *gorm.DB {
		return db.Preload(clause.Associations)
	})
}

func (d *Fire) CrudCount(model interface{}) (count int64, err error) {
	err = d.Model(model).Where(model).Count(&count).Error

	return
}

func (d *Fire) CrudSum(model interface{}, column string) (sum float32, err error) {
	err = d.Model(model).Where(model).Clauses(clause.Select{
		Expression: SUM{
			Column: clause.Column{
				Name: column,
			},
		},
	}).Find(&sum).Error

	return
}

type M map[string]interface{}

// CrudUpdate
// updates support (M or map[string]interface{}) and struct
// support gorm.Db.Select() and gorm.Db.Omit()
// TODO: struct only update non-zero fields
func (d *Fire) CrudUpdate(model interface{}, updates ...interface{}) (err error) {
	tx := d.Model(model).Where(model)
	if e, ok := tx.Statement.Clauses["WHERE"].Expression.(clause.Where); !ok || len(e.Exprs) <= 0 {
		return gorm.ErrMissingWhereClause
	}

	if len(updates) == 1 {
		tx = tx.Updates(updates[0])
	} else {
		// toMap
		resultMap := make(map[string]interface{})
		for _, item := range updates {
			switch reflect.TypeOf(item).Kind() {
			case reflect.Struct:
				if s, err := schema.Parse(item, &sync.Map{}, schema.NamingStrategy{}); err == nil {
					selectAll := arrayconv.StringIn("*", d.Statement.Selects)
					for _, field := range s.Fields {
						v, zero := field.ValueOf(d.Statement.Context, reflect.ValueOf(item))
						if selectAll || arrayconv.StringIn(field.DBName, d.Statement.Selects) || !zero {
							resultMap[field.DBName] = v
						}
					}
				}
			case reflect.Map:
				var m map[string]interface{}
				switch t := item.(type) {
				case M:
					m = t
				case map[string]interface{}:
					m = t
				}
				mergo.Merge(&resultMap, m)
			default:
				return gorm.ErrInvalidValue
			}
		}
		if len(resultMap) <= 0 {
			return errors.New("updates conditions required")
		}

		tx = tx.Updates(resultMap)
	}

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
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

	return NewInstance(d.Limit(1)).CrudUpdate(model, updates...)
}

func (d *Fire) CrudDelete(model interface{}) (err error) {
	// WARNING: When there is no condition WHERE, error is gorm.ErrMissingWhereClause("WHERE conditions required"), so you can safely delete it.
	err = d.Where(model).Delete(model).Error

	return
}
