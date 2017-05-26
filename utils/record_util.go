package utils

import (
	"log"
	"reflect"
)

func GetRecordColumns(record interface{}) (columns []string) {

	structure := reflect.ValueOf(record).Elem()
	structureType := reflect.TypeOf(record).Elem()

	for i := 0; i < structureType.NumField(); i++ {
		structureField := structureType.Field(i)
		field := structure.FieldByName(structureField.Name)
		if field.IsValid() && field.CanSet() {
			// get json tag to map from db
			column := structureType.Field(i).Tag.Get("json")
			if column != "" {
				columns = append(columns, column)
			}
		}
	}
	return
}

func SetData(record interface{}, data map[string]*[]byte) {
	s := reflect.ValueOf(record).Elem()
	typ := reflect.TypeOf(record).Elem()

	if typ.Kind() == reflect.Struct {
		for i := 0; i < typ.NumField(); i++ {

			structField := typ.Field(i)

			f := s.FieldByName(structField.Name)
			if f.IsValid() && f.CanSet() {

				// get json tag to map from db
				column := typ.Field(i).Tag.Get("json")

				// get the value returned from the database with the json field

				value := data[column]

				if value != nil {
					switch f.Kind() {
					case reflect.String:

						f.SetString(ParseByteArray(value, reflect.String).(string))
					case reflect.Int:
						f.SetInt(int64(ParseByteArray(value, reflect.Int).(int)))
					case reflect.Bool:
						x := ParseByteArray(value, reflect.Int).(int)
						f.SetBool(x == 1)
					default:
						f.SetBytes(*value)
					}
				}
			}
		}
	} else {
		log.Fatal("cannot set data of non-struct")
	}
}
