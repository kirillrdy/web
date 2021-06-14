package db

import (
	"fmt"
	"strings"
)

type Table string
type Column struct {
	Name        string
	Table       Table
	DisplayName string
}

func (column Column) fullName() string {
	//TODO for now hack here is easier
	if strings.HasPrefix(column.Name, "st_astext") {
		return column.Name
	}
	return fmt.Sprintf("\"%s\".\"%s\"", column.Table, column.Name)
}

func (table Table) Column(name string) Column {
	return Column{Name: name, Table: table}
}

func (table Table) ColumnWithDisplayName(name string, displayName string) Column {
	return Column{Name: name, Table: table, DisplayName: displayName}
}