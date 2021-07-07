/**
* @Author: TheLife
* @Date: 2021/5/17 下午5:42
 */
package goqueryPix

import (
	"github.com/PuerkitoBio/goquery"
	"reflect"
	"strings"
)

// goquery 工具，用于根据dom节点内容，自动映射到变量里。

// selection 根据 list，自动设置值
// 适合静态列表（多个相同类型对象，列表字段位置顺序都是一致情况用）
func SelectionPixList(element *goquery.Selection, arr []*string, iterationFuc ...func(value string, selection *goquery.Selection, index int) string) {
	element.Each(func(i int, selection *goquery.Selection) {
		if len(arr) > i {
			res := TrimAll(selection.First().Text())
			if iterationFuc != nil {
				res = iterationFuc[0](res, selection, i)
			}
			rv := reflect.ValueOf(arr[i]).Elem()
			rv.Set(reflect.ValueOf(res))
		}
	})
}

type M map[string]*string

// selection 根据 map，自动设置值
// 适合动态列表（多个相同类型对象，列表字段位置顺序不一致情况用）
func SelectionPixMap(element *goquery.Selection, key string, value string, m M, iterationFuc ...func(value string, selection *goquery.Selection, mKey string) string) {
	res := make(map[string]*goquery.Selection)

	element.Each(func(i int, selection *goquery.Selection) {
		res[selection.Find(key).Text()] = selection
	})

	for mKey, _ := range m {
		for resKey, resSelection := range res {
			// vague
			if strings.HasPrefix(resKey, mKey) {
				res := TrimAll(res[resKey].Find(value).First().Text())
				if iterationFuc != nil {
					res = iterationFuc[0](res, resSelection, mKey)
				}
				rv := reflect.ValueOf(m[mKey]).Elem()
				rv.Set(reflect.ValueOf(res))

				break
			}
		}
	}
}


func TrimAll(s string, cutset ...string) string {
	if s == "" {
		return s
	}

	if len(cutset) == 0 {
		cutset = []string{"\n", "\t", " "}
	}

	for _, value := range cutset {
		s = strings.Trim(s, value)
	}

	return s
}