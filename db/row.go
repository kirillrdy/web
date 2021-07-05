package db

import (
	"fmt"
	"log"
	"time"
)

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
