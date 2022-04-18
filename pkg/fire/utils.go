package fire

import (
	"gorm.io/gorm/clause"
	"net/url"
	"strconv"
	"strings"
	"unicode"
)

func toCamel2Case(m url.Values) {
	for key, value := range m {
		if !strings.Contains(key, "_") {
			delete(m, key)
			m[Camel2Case(key)] = value
		}
	}
}
func Camel2Case(name string) string {
	buffer := strings.Builder{}
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.WriteString("_")
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

func If(isA bool, a, b interface{}) interface{} {
	if isA {
		return a
	}

	return b
}

func ParseDataType(data string) interface{} {
	float, err := strconv.ParseFloat(data, 64)
	if err != nil {
		// string
		return data
	}
	// 46 = byte('.')
	if strings.IndexByte(data, 46) != -1 {
		// float
		return float
	}

	// int
	return int64(float)
}

func ParseColumn(c interface{}) (col clause.Column) {
	switch v := c.(type) {
	case clause.Column:
		col = v
	case string:
		col.Table = clause.CurrentTable
		col.Name = v
	}

	return col
}
