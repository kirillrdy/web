package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type InsertQuery struct {
	table   Table
	columns []Column
	values  []interface{}
}

func Insert() InsertQuery {
	return InsertQuery{}
}

func (insert InsertQuery) Into(table Table, columns ...Column) InsertQuery {
	insert.table = table
	insert.columns = columns
	return insert
}
func (insert InsertQuery) Values(values ...interface{}) InsertQuery {
	insert.values = values
	return insert
}
func (insert InsertQuery) Execute(db *sql.DB) (int64, error) {

	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("INSERT into %s", insert.table))

	var columns []string
	for _, column := range insert.columns {
		columns = append(columns, string(column.Name))
	}

	builder.WriteString("(")
	builder.WriteString(strings.Join(columns, ","))
	builder.WriteString(")")

	var values []string
	for index := range insert.values {
		values = append(values, fmt.Sprintf("$%d", index+1))
	}
	builder.WriteString(" VALUES(")
	builder.WriteString(strings.Join(values, ","))
	builder.WriteString(") RETURNING id") //TODO not all tables have id

	log.Print(builder.String())
	result, err := db.Exec(builder.String(), insert.values...)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()

}
