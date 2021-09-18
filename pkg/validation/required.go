package validation

import "reflect"

var Required = required{}

type required struct {
	message string
}

func (r required) Validate(value interface{}) bool {
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.String:
		return val.Len() == 0
	case reflect.Bool:
		return !val.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return val.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return val.IsNil()
	}
	return reflect.DeepEqual(val.Interface(), reflect.Zero(val.Type()).Interface())
}

func (r *required) SetMessage(message string) ValidateRule {
	r.message = message
	return r
}

func (r *required) Message() string {
	return r.message
}
