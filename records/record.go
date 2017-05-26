package records

import (
	"encoding/json"
	"fmt"
	"github.com/crwgregory/golang-api-skeleton/connection"
	"github.com/crwgregory/golang-api-skeleton/errors"
	"github.com/crwgregory/golang-api-skeleton/utils"
	"net/http"
	"reflect"
	"runtime/debug"
	"time"
)

type Record struct {
	connection connection.ConnectionInterface
}

type JSONData interface{}
type JSONObject map[string]interface{}
type JSONArray []JSON
type JSON map[string]interface{}
type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%d", time.Time(t).Unix())
	return []byte(stamp), nil
}

func GetRecordColumns(record interface{}) (columns []string) {

	structure := reflect.ValueOf(record).Elem()
	structureType := reflect.TypeOf(record).Elem()

	for i := 0; i < structureType.NumField(); i++ {
		structureField := structureType.Field(i)
		field := structure.FieldByName(structureField.Name)
		if field.IsValid() && field.CanSet() {
			// get json tag to map from db
			column := structureType.Field(i).Tag.Get("json")
			table := structureType.Field(i).Tag.Get("table") // if there is a table field set, this column does not belong to structs main table
			if column != "" && table == "" {
				columns = append(columns, column)
			}
		}
	}
	return
}

func GetFieldDataType(record interface{}, jsonFieldName string) (kind reflect.Kind, err error) {
	s := reflect.ValueOf(record).Elem()
	typ := reflect.TypeOf(record).Elem()
	if typ.Kind() == reflect.Struct {
		for i := 0; i < typ.NumField(); i++ {
			structField := typ.Field(i)
			f := s.FieldByName(structField.Name)
			if f.IsValid() && f.CanSet() {
				column := typ.Field(i).Tag.Get("json")
				if column == jsonFieldName {
					return f.Kind(), nil
				}
			}
		}
	} else {
		dErr := new(errors.DataConversionError)
		dErr.Message = "Cannot reflect of none struct type"
		dErr.Stack = debug.Stack()
		dErr.StatusCode = http.StatusUnprocessableEntity
		return reflect.Interface, dErr
	}
	return
}

func SetField(record interface{}, jsonFieldName string, data interface{}) error {
	field := getStructField(record, jsonFieldName)
	if &field == nil {
		return &errors.DataNotFoundError{}
	}
	value := reflect.ValueOf(data)
	field.Set(value)
	return nil
}

func getStructField(s interface{}, jsonFieldName string) *reflect.Value {
	st := reflect.ValueOf(s).Elem()
	typ := reflect.TypeOf(s).Elem()
	if typ.Kind() == reflect.Struct {
		for i := 0; i < typ.NumField(); i++ {
			structField := typ.Field(i)
			f := st.FieldByName(structField.Name)
			if f.IsValid() && f.CanSet() {
				column := typ.Field(i).Tag.Get("json")
				if column == jsonFieldName {
					return &f
				}
			}
		}
	}
	return nil
}

func SetData(record interface{}, data map[string]*[]byte) error {
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
						f.SetString(utils.ParseByteArray(value, reflect.String).(string))
					case reflect.Int:
						f.SetInt(int64(utils.ParseByteArray(value, reflect.Int).(int)))
					case reflect.Bool:
						x := utils.ParseByteArray(value, reflect.Int).(int)
						f.SetBool(x == 1)
					case reflect.Slice:

						_, ok := f.Interface().(JSONArray)
						if ok {
							var data JSONArray
							json.Unmarshal(*value, &data)
							f.Set(reflect.ValueOf(data))
						}
					case reflect.Struct:
						// can we set a struct?
					case reflect.Interface:

						var data JSONData
						if err := json.Unmarshal(*value, &data); err != nil {
							return err
						}
						f.Set(reflect.ValueOf(data))

					default:
						f.SetBytes(*value)
					}
				}
			}
		}
	} else {
		dErr := new(errors.DataConversionError)
		dErr.Message = "Cannot reflect of none struct type"
		dErr.Stack = debug.Stack()
		dErr.StatusCode = http.StatusUnprocessableEntity
		return dErr
	}
	return nil
}
