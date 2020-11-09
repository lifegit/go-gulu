/**
* @Author: TheLife
* @Date: 2020-11-8 6:23 下午
 */
package where

import "fmt"

type CompareType string

const (
	CompareEqual      CompareType = "="
	CompareAboutEqual CompareType = ">="
	CompareAbout      CompareType = ">"
	CompareLessEqual  CompareType = "<="
	CompareLess       CompareType = "<"
)

// field Type ?
type Compare struct {
	Field string
	Type  CompareType
	Text  interface{}
}

func (c *Compare) String() (query string, args interface{}) {
	return fmt.Sprintf("`%s` %s ?", c.Field, c.Type), c.Text
}
