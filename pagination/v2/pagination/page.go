/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package pagination

type Page struct {
	Total int64 `json:"total"`
	Size  int `json:"size"`
}
