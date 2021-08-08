package db

import (
	"fmt"
	"strings"
)

type Table string

var columns = make(map[Table][]Column)

func (table Table) Columns() []Column {
	return columns[table]
}

func (table Table) PrimaryKey() Column {
	//TODO for now hardcoded
	return Column{Name: "id", Table: table}
}

type Column struct {
	Table Table
	Name  string
}

func (column Column) PrimaryKey() bool {
	return column.Name == "id"
}

func (column Column) FullName() string {
	//TODO for now hack here is easier
	if strings.HasPrefix(column.Name, "st_astext") {
		return column.Name
	}
	return fmt.Sprintf("%s.%s", column.Table, column.Name)
}

func (table Table) Column(name string) Column {
	column := Column{Name: name, Table: table}
	columns[table] = append(columns[table], column)
	return column
}

func (column Column) Eq(value interface{}) WhereCondition {
	return WhereCondition{fragment: column.FullName() + " = ", arg: value}
}
