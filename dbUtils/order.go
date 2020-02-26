/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package dbUtils

const OrderAsc = "asc"
const OrderDesc = "desc"

type Order struct {
	Field string `form:"name" binding:"required"`
	Type  string `binding:"required,eq=asc|eq=desc"`
}
