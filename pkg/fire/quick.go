package fire

import (
	"fmt"
)

func (d *Fire) SumLabel(column ...string) (lab string) {
	if d.Config.Name() == "postgres" {
		lab = fmt.Sprintf("COALESCE(SUM(%s),0)", FormatColumn(column...))
	} else {
		lab = fmt.Sprintf("IFNULL(SUM(%s),0)", FormatColumn(column...))
	}
	return
}

func CountLabel(column ...string) (lab string) {
	c := If(column == nil, []string{ColumnAll}, column).([]string)[0]

	return fmt.Sprintf("COUNT(%s)", FormatColumn(c))
}
func UnnestLabel(column ...string) (lab string) {
	return fmt.Sprintf("UNNEST(%s)", FormatColumn(column...))
}

func DistinctOnLabel(column ...string) (lab string) {
	lab = FormatColumn(column...)
	return fmt.Sprintf("DISTINCT ON(%s) %s", lab, lab)
}

func AsLabel(column, asColumn string) (lab string) {
	return fmt.Sprintf("%s AS %s", column, FormatColumn(asColumn))
}
