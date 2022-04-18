package fire

import (
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm/clause"
	"reflect"
)

// === SELECT ===

func Select(exps ...interface{}) clause.Select {
	var expressions []clause.Expression

	for _, exp := range exps {
		switch v := exp.(type) {
		case string:
			expressions = append(expressions, clause.Expr{SQL: "?", Vars: []interface{}{ParseColumn(v)}})
		case clause.Column:
			expressions = append(expressions, clause.Expr{SQL: "?", Vars: []interface{}{v}})
		case clause.Expression:
			expressions = append(expressions, v)
		}
	}

	return clause.Select{
		Expression: clause.CommaExpression{Exprs: expressions},
	}
}

// === where ===

type CompareType string

const (
	CompareEq  CompareType = "="
	CompareGte CompareType = ">="
	CompareGt  CompareType = ">"
	CompareLte CompareType = "<="
	CompareLt  CompareType = "<"
)

// WhereCompare
// column CompareEqual ? # column = ?
func WhereCompare(column interface{}, value interface{}, compare ...CompareType) clause.Expression {
	c := If(compare != nil, compare, []CompareType{CompareEq}).([]CompareType)[0]

	col := ParseColumn(column)
	var exp clause.Expression
	switch c {
	case CompareEq:
		exp = clause.Eq{Column: col, Value: value}
	case CompareGte:
		exp = clause.Gte{Column: col, Value: value}
	case CompareGt:
		exp = clause.Gt{Column: col, Value: value}
	case CompareLte:
		exp = clause.Lte{Column: col, Value: value}
	case CompareLt:
		exp = clause.Lt{Column: col, Value: value}
	}

	return exp
}

func WhereJsonEq(column interface{}, value interface{}, key ...string) clause.Expression {
	return datatypes.JSONQuery(ParseColumn(column)).Equals(value, key...)
}

func WhereJsonHas(column interface{}, key ...string) clause.Expression {
	return datatypes.JSONQuery(ParseColumn(column)).HasKey(key...)
}

// WhereIn
// column IN(?)
// column NOT IN(?)
func WhereIn(column interface{}, value interface{}, isNot ...bool) clause.Expression {
	not := If(isNot != nil, isNot, []bool{false}).([]bool)[0]
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Slice:
		l, values := v.Len(), []interface{}{}
		for i := 0; i < l; i++ {
			s := v.Index(i)
			values = append(values, s.Interface())
		}
		var exp clause.Expression = clause.IN{
			Column: ParseColumn(column),
			Values: values,
		}
		if not {
			exp = clause.Not(exp)
		}
		return exp

	default:
		exp := clause.NamedExpr{
			fmt.Sprintf("? %s (?)", If(not, "NOT IN", "IN")), []interface{}{ParseColumn(column), value}}
		return exp
	}
}

type LikeType string

const (
	LikeFullText LikeType = "FullText"
	LikePrefix   LikeType = "Begin"
	LikeSuffix   LikeType = "End"
)

// WhereLike
// column LIKE %?%
func WhereLike(column interface{}, value interface{}, like ...LikeType) clause.Expression {
	c := If(like != nil, like, []LikeType{LikeFullText}).([]LikeType)[0]

	var v string
	switch c {
	case LikeFullText:
		v = fmt.Sprintf("%%%s%%", value)
	case LikePrefix:
		v = fmt.Sprintf("%%%s", value)
	case LikeSuffix:
		v = fmt.Sprintf("%s%%", value)
	default:
		v = fmt.Sprint(value)
	}
	return clause.Like{
		Column: ParseColumn(column),
		Value:  v,
	}
}

// WhereRange
// column >= start ANd column <= end
func WhereRange(column interface{}, start interface{}, end interface{}) clause.Expression {
	col := ParseColumn(column)
	return clause.Where{Exprs: []clause.Expression{
		clause.Gte{Column: col, Value: start},
		clause.Lte{Column: col, Value: end},
	}}
}

// === order ===

type OrderType string

const (
	OrderAsc  OrderType = "asc"
	OrderDesc OrderType = "desc"
)

func Order(column interface{}, order OrderType) clause.OrderByColumn {
	return clause.OrderByColumn{
		Column:  ParseColumn(column),
		Desc:    order == OrderDesc,
		Reorder: true,
	}
}

// === update ===

// UpdateArithmetic
// field = field ArithmeticType Number # field = field + 1
func UpdateArithmetic(column string, value interface{}, art ArithmeticType) (m map[string]interface{}) {
	return map[string]interface{}{
		column: SetArithmetic{Type: art, Column: clause.Column{Name: column}, Value: value},
	}
}
