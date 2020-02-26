/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package pagination

import "go-gulu/dbUtils"

type Param struct {
	Page     uint                    `form:"page" binding:"required,min=1"`
	Filtered *map[string]interface{} `form:"filtered"`
	Sorted   *dbUtils.Order          `form:"sorted"`
}