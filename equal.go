package db

import "fmt"

// Equal TODO: NEEDS COMMENT INFO
func Equal(field string, value interface{}) Operator {
	return &equalOperator{field, value}
}

type equalOperator struct {
	field string
	value interface{}
}

// Make TODO: NEEDS COMMENT INFO
func (o *equalOperator) Make() (string, []interface{}) {
	return fmt.Sprintf("`%s` = ?", o.field), []interface{}{o.value}
}
