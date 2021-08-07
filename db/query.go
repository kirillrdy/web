package db

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type WhereCondition struct {
	fragment string
	arg      interface{}
}

type Query struct {
	selectColumns []Column
	fromTable     Table
	joins         []struct {
		table       Table
		leftColumn  Column
		rightColumn Column
	}
	whereConditions []WhereCondition
	limit           *int
}

func (query Query) Join(table Table, leftColumn, rightColumn Column) Query {
	query.joins = append(query.joins, struct {
		table       Table
		leftColumn  Column
		rightColumn Column
	}{table: table, leftColumn: leftColumn, rightColumn: rightColumn})
	return query
}

func (query Query) Where(condition WhereCondition) Query {
	query.whereConditions = append(query.whereConditions, condition)
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

func (query Query) ExecuteOne(db *sql.DB) (Row, error) {
	rows := query.Limit(1).Execute(db)
	if len(rows) == 0 {
		//TODO better errors and use conctant errors
		return Row{}, fmt.Errorf("didnt find anything")
	}
	return rows[0], nil
}

func (query Query) Execute(db *sql.DB) []Row {
	var builder strings.Builder
	builder.WriteString("SELECT ")
	for i, column := range query.selectColumns {
		if i == len(query.selectColumns)-1 {
			builder.WriteString(column.FullName())
		} else {
			builder.WriteString(column.FullName() + ",")
		}
	}
	builder.WriteString(" FROM ")
	builder.WriteString(string(query.fromTable))

	for _, join := range query.joins {
		builder.WriteString(fmt.Sprintf(" JOIN %s ON %s = %s ", join.table, join.leftColumn.FullName(), join.rightColumn.FullName()))
	}

	if len(query.whereConditions) != 0 {
		builder.WriteString(" WHERE ")
	}
	var args []interface{}
	for index, condition := range query.whereConditions {
		if index == len(query.whereConditions)-1 {
			builder.WriteString(condition.fragment + fmt.Sprintf(" $%d ", index+1))
		} else {
			builder.WriteString(condition.fragment + fmt.Sprintf(" $%d AND ", index+1))
		}
		args = append(args, condition.arg)
	}

	// Default limit
	limit := 10
	if query.limit != nil {
		limit = *query.limit
	}

	builder.WriteString(fmt.Sprintf(" LIMIT %d", limit))
	log.Printf("QUERY: %s", builder.String())

	rows, err := db.Query(builder.String(), args...)
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
