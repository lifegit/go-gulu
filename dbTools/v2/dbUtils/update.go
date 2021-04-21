/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package dbUtils

import "go-gulu/dbTools/v2/tool/update"

type DbUpdates struct {
	Arithmetic []update.Arithmetic
	Set        []update.Set
}

func (d *DbUtils) Update(w interface{}) *DbUtils {
	if d.Updates == nil {
		d.Updates = &DbUpdates{}
	}

	switch t := w.(type) {
	case update.Arithmetic:
		d.Updates.Arithmetic = append(d.Updates.Arithmetic, t)
	case update.Set:
		d.Updates.Set = append(d.Updates.Set, t)
	}

	return d
}

func (d *DbUtils) GetUpdate() (m *map[string]interface{}) {
	if d == nil || d.Updates == nil {
		return nil
	}
	mp := make(map[string]interface{})
	for _, value := range d.Updates.Arithmetic {
		mp[value.Field] = value.String()
	}
	for _, value := range d.Updates.Set {
		mp[value.Field] = value.String()
	}

	return &mp
}
