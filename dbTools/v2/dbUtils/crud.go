/**
* @Author: TheLife
* @Date: 2021/5/27 下午5:08
 */
package dbUtils

import (
	"errors"
	"fmt"
)

func (d *DbUtils) IsExists(model interface{}) bool {
	var s interface{}
	res := d.Model(model).Where(model).Select("1").Take(&s)

	return res.Error == nil
}

func (d *DbUtils) CrudOne(model interface{}, callData interface{}) (err error) {
	err = d.Model(model).Where(model).Take(callData).Error

	return
}

func (d *DbUtils) CrudAll(model interface{}, callListData interface{}) (err error) {
	err = d.Model(model).Where(model).Find(callListData).Error

	return
}

func (d *DbUtils) CrudAllPage(model interface{}, callListData interface{}, page PageParam, allow Allow) (total int64, err error) {
	// SELECT * FROM `user` WHERE (`time` >= 1622173028456.000000 AND `time` <= 1622259431456.000000) AND `name` LIKE '%张明%' AND `age` = 18.000000 ORDER BY id asc

	tx := d.Model(model).Where(model)
	tx = allow.AllowParams(page.Params, d).DB
	//tx = allow.Allow(page.Params, page.Sort, d).DB

	tx.Find(callListData) // .Offset(20).Limit(20)

	tx.Count(&total)

	////return
	//
	//tx.Count(&count)
	//
	////err = tx.Count(&count).Error
	//if count >= 0 {
	//	err =
	//}

	return
}

func (d *DbUtils) CrudCount(model interface{}) (count int64, err error) {
	err = d.Model(model).Where(model).Count(&count).Error

	return
}

func (d *DbUtils) CrudSum(model interface{}, column string) (sum float32, err error) {
	err = d.Model(model).Where(model).Select(fmt.Sprintf("IFNULL(SUM(`%s`),0)", column)).Row().Scan(&sum)

	return
}

func (d *DbUtils) CrudUpdate(model interface{}, updates interface{}, updateOne bool) (err error) {
	tx := d.Model(model).Where(model)

	if updateOne {
		tx = tx.Limit(1)
	}

	tx = tx.Updates(updates)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected <= 0 {
		return errors.New("resource is not found")
	}
	return
}

func (d *DbUtils) CrudDelete(model interface{}) (err error) {
	//WARNING When delete a record, you need to ensure it’s primary field has value, and GORM will use the primary key to delete the record, if primary field’s blank, GORM will delete all records for the model
	//primary key must be not zero value
	err = d.Where(model).Delete(model).Error

	return
}
