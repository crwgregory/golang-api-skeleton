package utils

import (
	"reflect"
)

func ArrayContains(array interface{}, contains interface{}, kind reflect.Kind) (has bool, err error) {
	switch kind {
	case reflect.String:
		a, ok := array.([]string)
		if !ok {
			return false, returnConversionError("string")
		}
		c, ok := contains.(string)
		if !ok {
			return false, returnConversionError("string")
		}
		for _, v := range a {
			if v == c {
				return true, nil
			}
		}
	}
	return
}

func returnConversionError(conversion string) error {
	dErr := new(errors.DataConversionError)
	dErr.Message = "cannot convert to " + conversion
	return dErr
}
