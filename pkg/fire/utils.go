package fire

import (
	"net/url"
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
