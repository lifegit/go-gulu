/**
* @Author: TheLife
* @Date: 2020-4-24 1:38 下午
 */
package dbUtils

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

func (d *DbUtils) CrudOne(fields []string, where interface{}, one interface{}, defaultDb *gorm.DB) error {
	tx := InitDb(d, defaultDb)
	tx = tx.Where(where)
	tx = d.GetWhere(tx)
	tx = d.GetOrder(tx)
	if fields != nil {
		tx = tx.Select(fields)
	}

	res := tx.Debug().Take(one)
	if res.Error != nil {
		return res.Error
	}
	if res.RecordNotFound() {
		return errors.New("resource is not found")
	}

	return nil
}

func (d *DbUtils) CrudAll(fields []string, where interface{}, list interface{}, defaultDb *gorm.DB) error {
	tx := InitDb(d, defaultDb)
	tx = tx.Where(where)
	tx = d.GetWhere(tx)
	tx = d.GetOrder(tx)
	if fields != nil {
		tx = tx.Select(fields)
	}

	err := tx.Debug().Find(list).Error
	return err
}

func (d *DbUtils) CrudAllPage(fields []string, where interface{}, list interface{}, limit uint, defaultDb *gorm.DB) (count uint, err error) {
	tx := InitDb(d, defaultDb)
	tx = tx.Model(where).Where(where)
	tx = d.GetWhere(tx)
	tx = d.GetOffSet(tx)
	tx = d.GetOrder(tx)
	tx = d.GetJoin(tx)

	if fields != nil {
		tx = tx.Select(fields)
	}

	err = tx.Count(&count).Error
	if count > 0 {
		err = tx.Debug().Limit(limit).Find(list).Error
	}

	return count, err
}

func (d *DbUtils) CrudCount(where interface{}, defaultDb *gorm.DB) (count uint, err error) {
	tx := InitDb(d, defaultDb)
	tx = tx.Model(where).Where(where)
	tx = d.GetWhere(tx)

	err = tx.Count(&count).Error
	return
}
func (d *DbUtils) CrudSum(field string, where interface{}, defaultDb *gorm.DB) (sum float32, err error) {
	tx := InitDb(d, defaultDb)
	tx = tx.Model(where).Where(where)
	tx = d.GetWhere(tx)

	row := tx.Select(fmt.Sprintf("IFNULL(SUM(`%s`),0)", field)).Row()
	err = row.Scan(&sum)

	return
}
func (d *DbUtils) CrudUpdate(updates map[string]interface{}, where interface{}, defaultDb *gorm.DB, limit1 bool) (err error) {
	tx := InitDb(d, defaultDb)
	tx = d.GetWhere(tx)
	if up := d.GetUpdate(); up != nil {
		for k, v := range *up {
			updates[k] = v
		}
	}

	tx = tx.Debug().Model(where)
	if limit1 {
		tx.Limit(1)
	}
	tx = tx.Updates(updates)
	if err = tx.Error; err != nil {
		return
	}
	if tx.RowsAffected <= 0 {
		return errors.New("resource is not found")
	}
	return nil
}

func (d *DbUtils) CrudDelete(where interface{}, defaultDb *gorm.DB) error {
	//WARNING When delete a record, you need to ensure it’s primary field has value, and GORM will use the primary key to delete the record, if primary field’s blank, GORM will delete all records for the model
	//primary key must be not zero value
	tx := InitDb(d, defaultDb)
	tx = tx.Where(where)
	tx = d.GetWhere(tx)

	tx = tx.Debug().Delete(where)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
