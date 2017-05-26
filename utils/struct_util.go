package utils

import (
	"reflect"
)

func GetFieldValue(s interface{}, fieldName string) interface{} {
	st := reflect.ValueOf(s).Elem()
	typ := reflect.TypeOf(s).Elem()
	for i := 0; i < typ.NumField(); i++ {
		structField := typ.Field(i)
		if structField.Name == fieldName {
			f := st.FieldByName(structField.Name)

			if f.IsValid() && f.CanSet() && f.CanAddr() && !f.IsNil() {
				return f.Interface()
			}
		}

	}
	return nil
}
