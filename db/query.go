package db

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

var DB *sql.DB

type Query struct {
	selectColumns []Column
	fromTable     Table
	joins         []struct {
		table       Table
		leftColumn  Column
		rightColumn Column
	}
	limit *int
}

func (query Query) Join(table Table, leftColumn, rightColumn Column) Query {
	query.joins = append(query.joins, struct {
		table       Table
		leftColumn  Column
		rightColumn Column
	}{table: table, leftColumn: leftColumn, rightColumn: rightColumn})
	return query
}

func Select(columns ...Column) Query {
	return Query{selectColumns: columns}
}

func (query Query) Limit(limit int) Query {
	query.limit = &limit
	return query
}

func (query Query) From(table Table) Query {
	query.fromTable = table
	return query
}

//TODO remove this, sub packages should return erros instead of panic
func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

type Row struct {
	columnOrder map[Column]int
	values      []interface{}
}

func (row Row) String(column Column) string {
	order, ok := row.columnOrder[column]
	if !ok {
		log.Printf("WARNING: column not found in maping %v", column)
	}
	return fmt.Sprintf("%v", row.values[order])
}

func (row Row) GetString(column Column) string {
	index := row.columnOrder[column]
	value, ok := row.values[index].(string)
	if !ok {
		log.Printf("WARNING: %s is not a string", column.fullName())
	}
	return value
}

func (query Query) Execute() []Row {
	var builder strings.Builder
	builder.WriteString("SELECT ")
	for i, column := range query.selectColumns {
		if i == len(query.selectColumns)-1 {
			builder.WriteString(column.fullName())
		} else {
			builder.WriteString(column.fullName() + ",")
		}
	}
	builder.WriteString(" FROM ")
	builder.WriteString(string(query.fromTable))

	for _, join := range query.joins {
		builder.WriteString(fmt.Sprintf(" JOIN %s ON %s = %s ", join.table, join.leftColumn.fullName(), join.rightColumn.fullName()))
	}

	// Default limit
	limit := 10
	if query.limit != nil {
		limit = *query.limit
	}

	builder.WriteString(fmt.Sprintf(" LIMIT %d", limit))
	log.Printf("QUERY: %s", builder.String())

	rows, err := DB.Query(builder.String())
	check(err)

	defer rows.Close()
	var results []Row

	columMap := make(map[Column]int)
	for i := range query.selectColumns {
		columMap[query.selectColumns[i]] = i
	}

	for rows.Next() {
		result := make([]interface{}, len(query.selectColumns))
		for i := range result {
			var empty interface{}
			result[i] = &empty
		}
		err := rows.Scan(result...)
		values := make([]interface{}, len(result))
		for i := range result {
			values[i] = reflect.ValueOf(result[i]).Elem().Interface()
		}
		results = append(results, Row{columnOrder: columMap, values: values})
		check(err)
	}

	return results
}
