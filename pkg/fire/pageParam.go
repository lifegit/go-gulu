// Package fire /***/
// 定义了分页、筛选、排序等的数据模型。支持 antd v5

package fire

// PageResult 分页结果
type PageResult struct {
	Page
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

func (p *PageResult) SetData(d interface{}) {
	p.Data = d
}

func (p *PageResult) Init(page ...Page) {
	if page != nil {
		p.Page = page[0]
	}

	if p.PageSize <= 0 {
		p.PageSize = DefaultPageSize
	}
	if p.Page.Current <= 0 {
		p.Current = 1
	}
}

type Page struct {
	Current  int `json:"current" form:"current"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

func (p *Page) GetOffset() int {
	return p.PageSize * (p.Current - 1)
}

const DefaultPageSize = 20

func (p *Page) DefaultSize(pageSize int) Page {
	if p.PageSize <= 0 {
		p.PageSize = pageSize
	}

	return *p
}

func (p *Page) SetSize(pageSize int) *Page {
	p.PageSize = pageSize

	return p
}

// SinglePageResult 单页结果
type SinglePageResult struct {
	Data interface{} `json:"data"`
}
