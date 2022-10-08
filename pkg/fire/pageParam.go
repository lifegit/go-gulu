// Package fire /***/
// 定义了分页、筛选、排序等的数据模型。支持 antd v5

package fire

import (
	"net/url"
	"strconv"
)

var DefaultPageSize = 5

// PageResult 分页结果
// https://gorm.io/docs/scopes.html#Pagination
type PageResult struct {
	Page
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

func (p *PageResult) SetData(d interface{}) {
	p.Data = d
}

func (p *PageResult) Init(page ...interface{}) {
	// default
	p.Page = Page{
		Current:  1,
		PageSize: DefaultPageSize,
	}

	if page != nil {
		switch v := page[0].(type) {
		case Page:
			p.Page = v
		case *url.URL:
			query := v.Query()
			current, _ := strconv.Atoi(query.Get("current"))
			if current <= 0 {
				current = 1
			}
			p.Current = current

			pageSize, _ := strconv.Atoi(query.Get("pageSize"))
			switch {
			case pageSize > 50:
				pageSize = 50
			case pageSize <= 0:
				pageSize = DefaultPageSize
			}
			p.PageSize = pageSize
		}
	}
}

type Page struct {
	Current  int `json:"current" form:"current"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

func (p *Page) GetOffset() int {
	return p.PageSize * (p.Current - 1)
}

// SinglePageResult 单页结果
type SinglePageResult struct {
	Data interface{} `json:"data"`
}
