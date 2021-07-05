/**
* @Author: TheLife
* @Date: 2021/5/27 下午4:34
 */
package fire

// 定义了分页、筛选、排序等的数据模型。支持 antd

// 筛选、排序参数
type Param struct {
	Params Params `form:"params" json:"params"`
	Sort   Sort   `form:"sort" json:"sort" binding:"omitempty,eq=1,dive,keys,required,endkeys,eq=ascend|eq=descend"`
}

// 分页参数
type PageParam struct {
	Page
	Param
}

// 分页结果
type PageResult struct {
	Page
	Total int64 `json:"total"`
}

func (p *PageResult) Init(page ...Page) {
	if page != nil {
		p.Page = page[0]
	}

	if p.PageSize <= 0 {
		p.PageSize = 20
	}
	if p.Page.Current <= 0 {
		p.Current = 1
	}
}

// 页
type Page struct {
	Current  int `json:"current" form:"current"`
	PageSize int `json:"page_size" form:"page_size"`
}

func (p *Page) GetOffset() int {
	return p.PageSize * (p.Current - 1)
}
