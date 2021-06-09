/**
* @Author: TheLife
* @Date: 2021/5/27 下午4:34
 */
package dbUtils

type PageParam struct {
	Page
	Param
}

type Param struct {
	Params Params `form:"params" json:"params"`
	Sort   Sort   `form:"sort" json:"sort" binding:"omitempty,eq=1,dive,keys,required,endkeys,eq=ascend|eq=descend"`
}

type Page struct {
	Current  int `json:"current" form:"current"`
	PageSize int `json:"page_size" form:"page_size"`
}

func (p *Page) GetOffset() int {
	return p.PageSize * (p.Current - 1)
}
