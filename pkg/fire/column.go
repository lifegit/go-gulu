package fire

import (
	"fmt"
	"strings"
)

type FormatColumnType string

const (
	FormatColumnBackQuote      FormatColumnType = "`"
	FormatColumnQuotationMarks FormatColumnType = `"`
)

var formatColumn = string(FormatColumnBackQuote)

func SetFormatColumnType(v FormatColumnType) {
	formatColumn = string(v)
}

func FormatColumn(column ...string) (res string) {
	var list []string
	for _, item := range column {
		list = append(list, strings.Split(item, ".")...)
	}

	for key, value := range list {
		if len(value) >= 1 {
			// first and last is not `
			if value == ColumnAll {
				value = ColumnAll
			} else if value[:1] != formatColumn && value[len(value)-1:] != formatColumn {
				value = fmt.Sprintf("%s%s%s", formatColumn, value, formatColumn)
			}
			res += value
			// isLast
			if key != len(list)-1 {
				res += "."
			}
		}
	}

	return
}

const ColumnAll = "*"

type Column struct {
	Table  string
	Column string
}

func (c *Column) String() string {
	return FormatColumn(c.Table, c.Column)
}
