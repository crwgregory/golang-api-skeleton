package utils

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func ParseByteArray(data *[]byte, convertTo reflect.Kind) interface{} {
	switch convertTo {
	case reflect.Int:
		value := fmt.Sprintf("%s", data)[1:]
		id, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		return id

	case reflect.String:
		return fmt.Sprintf("%s", data)[1:]
	}
	return data
}

func ScanRow(row *sql.Row, columns []string) (results map[string]*[]byte, err error) {
	results = make(map[string]*[]byte)
	var rowData []interface{}

	for i := 0; i < len(columns); i++ {
		var column []byte
		rowData = append(rowData, &column)
	}

	if err := row.Scan(rowData...); err != nil {
		return nil, err
	}

	for i, column := range rowData {
		results[columns[i]] = column.(*[]byte)
	}
	return
}

func ScanRows(columns []string, rows *sql.Rows) (results map[string]*[]byte) {
	results = make(map[string]*[]byte)
	var rowData []interface{}

	for i := 0; i < len(columns); i++ {
		var column []byte
		rowData = append(rowData, &column)
	}

	if err := rows.Scan(rowData...); err != nil {
		log.Fatal(err)
	}

	for i, column := range rowData {
		results[columns[i]] = column.(*[]byte)
	}
	return
}
