package remap

import (
	"reflect"
)

// OnFields maps fields with the same name from one
// struct to another.
//  OnFields(&from, &to)
func OnFields(from interface{}, to interface{}) {
	fromMap := make(map[string]reflect.Value)

	tFrom := reflect.TypeOf(from).Elem()
	vFrom := reflect.ValueOf(from).Elem()

	for i := 0; i < tFrom.NumField(); i++ {
		fld := tFrom.Field(i)
		fromMap[fld.Name] = vFrom.Field(i)
	}

	tTo := reflect.TypeOf(to).Elem()
	vTo := reflect.ValueOf(to).Elem()

	for i := 0; i < tTo.NumField(); i++ {
		field := vTo.Field(i)
		fieldName := tTo.Field(i).Name

		if val, ok := fromMap[fieldName]; ok {
			if field.Kind() == val.Kind() {
				field.Set(val)
			} else if field.Kind() == reflect.Ptr {
				tmp := reflect.New(field.Type().Elem())
				tmp.Elem().Set(val)
				field.Set(tmp)
			} else if val.Kind() == reflect.Ptr {
				field.Set(val.Elem())
			}
		}
	}
}
