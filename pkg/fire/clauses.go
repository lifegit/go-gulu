package fire

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

// COUNT for select
type COUNT struct {
	Column clause.Column
}

func (c COUNT) Build(builder clause.Builder) {
	alias := c.Column.Alias
	c.Column.Alias = ""

	builder.WriteString("COUNT")
	builder.WriteByte('(')
	if c.Column.Name == "" {
		builder.WriteByte('*')
	} else {
		builder.WriteQuoted(c.Column)
	}
	builder.WriteByte(')')

	if stm, b := builder.(*gorm.Statement); alias != "" && b {
		var buf strings.Builder
		stm.QuoteTo(&buf, clause.Column{
			Alias: alias,
		})
		builder.WriteString(buf.String()[2:])
	}
}

type SUM struct {
	Column clause.Column
}

func (s SUM) Build(builder clause.Builder) {
	stm, b := builder.(*gorm.Statement)
	if !b {
		return
	}
	alias := s.Column.Alias
	s.Column.Alias = ""

	if stm.Config.Name() == "postgres" {
		builder.WriteString("COALESCE(SUM(")
	} else {
		builder.WriteString("IFNULL(SUM(")
	}
	builder.WriteQuoted(s.Column)
	builder.WriteString("),0)")

	if alias != "" {
		var buf strings.Builder
		stm.QuoteTo(&buf, clause.Column{
			Alias: alias,
		})
		builder.WriteString(buf.String()[2:])
	}
}

// UNNEST for select
type UNNEST struct {
	Column clause.Column
}

func (u UNNEST) Build(builder clause.Builder) {
	alias := u.Column.Alias
	u.Column.Alias = ""

	builder.WriteString("UNNEST")
	builder.WriteByte('(')
	builder.WriteQuoted(u.Column)
	builder.WriteByte(')')

	if alias != "" {
		stm, b := builder.(*gorm.Statement)
		if !b {
			return
		}
		var buf strings.Builder
		stm.QuoteTo(&buf, clause.Column{
			Alias: alias,
		})
		builder.WriteString(buf.String()[2:])
	}
}

type ArithmeticType string

const (
	ArithmeticIncrease ArithmeticType = "+"
	ArithmeticReduce   ArithmeticType = "-"
	ArithmeticMultiply ArithmeticType = "*"
	ArithmeticExcept   ArithmeticType = "/"
)

// SetArithmetic for set
type SetArithmetic struct {
	Column clause.Column
	Type   ArithmeticType
	Value  interface{}
}

func (a SetArithmetic) Build(builder clause.Builder) {
	builder.WriteQuoted(a.Column)
	builder.WriteString(fmt.Sprintf(" %s ", a.Type))
	builder.AddVar(builder, a.Value)
}
