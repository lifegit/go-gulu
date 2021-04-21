/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package pagination

import (
	"go-gulu/dbTools/tool/order"
)

type Param struct {
	Page     int                    `form:"page" binding:"required,min=1"`
	Filtered *map[string]interface{} `form:"filtered"`
	Sorted   *order.Order          `form:"sorted"`
}


