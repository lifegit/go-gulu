/**
* @Author: TheLife
* @Date: 2021/5/26 下午11:21
 */
package dbUtils

type Allow struct {
	Where []string
	Like  []string
	Range []string
	In    []string
}

// allow
func (a *Allow) Allow(params Params, sort Sort, utils *DbUtils) (res *DbUtils) {
	a.AllowParams(params, utils)
	a.AllowSort(sort, utils)

	return utils
}

type Sort map[string]string

// allowSort
func (a *Allow) AllowSort(sort Sort, utils *DbUtils) *DbUtils {
	for column, value := range sort {
		utils.OrderByColumn(column, If(value == "ascend", OrderAsc, OrderDesc).(OrderType))
	}

	return utils
}

type Params map[string]interface{}

// allowParams
func (a *Allow) AllowParams(params Params, utils *DbUtils) *DbUtils {
	for column, value := range params {
		// Where
		for _, item := range a.Where {
			if item == column {
				switch data := value.(type) {
				case []interface{}:
					if len(data) >= 1 {
						utils.WhereCompare(column, data[0])
					}
				case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128:
					utils.WhereCompare(column, value)
				}
				continue
			}
		}

		// Like
		for _, item := range a.Like {
			if item == column {
				switch value.(type) {
				case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128:
					utils.WhereLike(column, value)
				}
				continue
			}
		}

		// Range
		for _, item := range a.Range {
			if item == column {
				switch data := value.(type) {
				case []interface{}:
					if len(data) >= 2 {
						utils.WhereRange(column, data[0], data[1])
					}
				}
				continue
			}
		}

		// In
		for _, item := range a.In {
			if item == column {
				switch value.(type) {
				case []interface{}:
					utils.WhereIn(column, value)
				}
				continue
			}
		}

	}

	return utils
}
