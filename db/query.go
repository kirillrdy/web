package db

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"
)

var DB *sql.DB

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

type Row struct {
	columnOrder map[Column]int
	values      []interface{}
}

func (row Row) String(column Column) string {
	value := row.GetValue(column)
	if value, ok := value.(time.Time); ok {
		//TODO i guess this really needs to be user timezone not server Local
		return value.Local().Format(time.RFC822)
	}
	return fmt.Sprintf("%v", value)
}

// These are designed to be used in the views, so they dont panic, or return errors
func (row Row) GetValue(column Column) interface{} {
	index, ok := row.columnOrder[column]
	if !ok {
		log.Printf("WARNING: %v is not found in %v", column, row.columnOrder)
	}
	return row.values[index]
}

func (row Row) GetInt64(column Column) int64 {
	interfaceValue := row.GetValue(column)
	value, ok := interfaceValue.(int64)
	if !ok {
		log.Printf("WARNING: %s is not an int but %T", column.FullName(), interfaceValue)
	}
	return value
}

func (row Row) GetString(column Column) string {
	value, ok := row.GetValue(column).(string)
	if !ok {
		log.Printf("WARNING: %s is not a string", column.FullName())
	}
	return value
}

func (query Query) ExecuteOne() (Row, error) {
	rows := query.Limit(1).Execute()
	if len(rows) == 0 {
		//TODO better errors and use conctant errors
		return Row{}, fmt.Errorf("didnt find anything")
	}
	return rows[0], nil
}

func (query Query) Execute() []Row {
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

	rows, err := DB.Query(builder.String(), args...)
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
