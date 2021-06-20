package db

import (
	"fmt"
	"log"
	"strings"
)

type UpdateQuery struct {
	table          Table
	whereCondition WhereCondition
	values         []struct {
		column Column
		value  interface{}
	}
}

func Update(table Table) UpdateQuery {
	return UpdateQuery{table: table}
}

func (query UpdateQuery) Where(condition WhereCondition) UpdateQuery {
	query.whereCondition = condition
	return query
}

func (query UpdateQuery) Set(column Column, value interface{}) UpdateQuery {
	query.values = append(query.values, struct {
		column Column
		value  interface{}
	}{column: column, value: value})
	return query
}

func (query UpdateQuery) Execute() error {
	var builder strings.Builder
	builder.WriteString("UPDATE ")
	builder.WriteString(string(query.table))
	builder.WriteString(" SET ")
	var args []interface{}
	for index, value := range query.values {

		builder.WriteString(value.column.Name)
		builder.WriteString(" = ")
		builder.WriteString(fmt.Sprintf(" $%d", index+1))

		if index != len(query.values)-1 {
			builder.WriteString(" , ")
		}
		args = append(args, value.value)
	}

	args = append(args, query.whereCondition.arg)

	builder.WriteString(" WHERE ")
	builder.WriteString(query.whereCondition.fragment)
	builder.WriteString(fmt.Sprintf("$%d", len(query.values)+1))

	log.Print(builder.String())
	result, err := DB.Exec(builder.String(), args...)

	if err != nil {
		return err
	}
	log.Print(result)
	return nil
}
