/**
* @Author: TheLife
* @Date: 2021/5/19 下午4:12
 */
package goqueryPix_test

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/lifegit/go-gulu/v2/pkg/spider/goqueryPix"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var doc = `
  <ul class="general-item-wrap"> 
   <li class="intro-item"> <span class="title">类型</span> <span class="content">出租   </span> </li> 
   <li class="intro-item"> <span class="title">名称</span> <span class="content">公寓 		</span> </li>
  </ul>
`

type Pix struct {
	Type string
	Name string
}

var success Pix

func init() {
	success = Pix{
		Type: "出租",
		Name: "公寓",
	}
}

func TestSelectionPixList(t *testing.T) {
	dom, _ := goquery.NewDocumentFromReader(strings.NewReader(doc))

	var p Pix
	goqueryPix.SelectionPixList(dom.Find("li span.content"), []*string{&p.Type, &p.Name},
		func(value string, selection *goquery.Selection, index int) string {
			return goqueryPix.TrimAll(value)
		},
	)

	assert.Equal(t, p, success)
}

func TestSelectionPixMap(t *testing.T) {
	dom, _ := goquery.NewDocumentFromReader(strings.NewReader(doc))

	var p Pix
	goqueryPix.SelectionPixMap(dom.Find("li"), "span.title", "span.content", goqueryPix.M{
		"类":  &p.Type, // vague
		"名称": &p.Name,
	}, func(value string, selection *goquery.Selection, key string) string {
		return goqueryPix.TrimAll(value)
	})

	assert.Equal(t, p, success)
}
